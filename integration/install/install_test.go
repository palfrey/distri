package install_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/distr1/distri/internal/env"
	"github.com/distr1/distri/pb"
	"github.com/golang/protobuf/proto"
)

const (
	systemd = "systemd-amd64-239"
	bash    = "bash-amd64-4.4.18"
)

func installFile(ctx context.Context, tmpdir string, pkg ...string) error {
	install := exec.Command("distri",
		append([]string{
			"install",
			"-root=" + tmpdir,
			"-repo=" + env.DefaultRepo,
		}, pkg...)...)
	install.Stderr = os.Stderr
	install.Stdout = os.Stdout
	if err := install.Run(); err != nil {
		return fmt.Errorf("%v: %v", install.Args, err)
	}
	return nil
}

func export(ctx context.Context, repo string) (addr string, cleanup func(), _ error) {
	export := exec.CommandContext(ctx, "distri",
		"-addrfd=3", // Go dup2()s ExtraFiles to 3 and onwards
		"export",
		"-listen=localhost:0",
		// Disable gzip: the gzipped.FileServer package is already tested, and
		// uncompressing these files makes the test run significantly slower.
		"-gzip=false",
		"-repo="+repo,
	)
	r, w, err := os.Pipe()
	export.Stderr = os.Stderr
	export.Stdout = os.Stdout
	export.ExtraFiles = []*os.File{w}
	if err := export.Start(); err != nil {
		return "", nil, fmt.Errorf("%v: %v", export.Args, err)
	}
	cleanup = func() {
		export.Process.Kill()
		export.Wait()
	}

	// Close the write end of the pipe in the parent process.
	if err := w.Close(); err != nil {
		return "", nil, err
	}

	// Read the listening address from the pipe. A successful read also serves
	// as readiness notification.
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil, err
	}
	return string(b), cleanup, nil
}

func installHTTP(ctx context.Context, tmpdir string, pkg ...string) error {
	addr, cleanup, err := export(ctx, env.DefaultRepo)
	if err != nil {
		return err
	}
	defer cleanup()

	install := exec.Command("distri",
		append([]string{
			"install",
			"-root=" + tmpdir,
			"-repo=http://" + addr,
		}, pkg...)...)
	install.Stderr = os.Stderr
	install.Stdout = os.Stdout
	if err := install.Run(); err != nil {
		return fmt.Errorf("%v: %v", install.Args, err)
	}
	return nil
}

func readMeta(fn string) (*pb.Meta, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var m pb.Meta
	if err := proto.UnmarshalText(string(b), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func installHTTPMultiple(ctx context.Context, tmpdir string, pkg ...string) error {
	// Create temporary repos which only hold one package (and its runtime
	// dependencies):
	addrs := make(map[string]string) // pkg → addr
	for _, pkg := range []string{systemd, bash} {
		rtmpdir, err := ioutil.TempDir("", "distritest")
		if err != nil {
			return err
		}
		defer os.RemoveAll(rtmpdir)
		meta, err := readMeta(filepath.Join(env.DefaultRepo, pkg+".meta.textproto"))
		if err != nil {
			return err
		}
		for _, dep := range append([]string{pkg}, meta.GetRuntimeDep()...) {
			cp := exec.Command("cp",
				filepath.Join(env.DefaultRepo, dep+".squashfs"),
				filepath.Join(env.DefaultRepo, dep+".meta.textproto"),
				rtmpdir)
			cp.Stderr = os.Stderr
			if err := cp.Run(); err != nil {
				return fmt.Errorf("%v: %v", cp.Args, err)
			}
		}
		addr, cleanup, err := export(ctx, rtmpdir)
		if err != nil {
			return err
		}
		defer cleanup()
		addrs[pkg] = addr
	}

	ctmpdir, err := ioutil.TempDir("", "distritest")
	if err != nil {
		return err
	}
	defer os.RemoveAll(ctmpdir)
	reposd := filepath.Join(ctmpdir, "repos.d")
	if err := os.Mkdir(reposd, 0755); err != nil {
		return err
	}
	for pkg, addr := range addrs {
		if err := ioutil.WriteFile(filepath.Join(reposd, pkg), []byte("http://"+addr), 0644); err != nil {
			return err
		}
	}
	install := exec.Command("distri",
		append([]string{
			"install",
			"-root=" + tmpdir,
		}, pkg...)...)
	install.Env = []string{"DISTRICFG=" + ctmpdir}
	install.Stderr = os.Stderr
	install.Stdout = os.Stdout
	if err := install.Run(); err != nil {
		return fmt.Errorf("%v: %v", install.Args, err)
	}
	return nil
}

func installHTTPMultipleVersions(ctx context.Context, tmpdir string, pkg ...string) error {
	// Create temporary repos which only hold one package (and its runtime
	// dependencies):
	addrs := make(map[string]string) // pkg → addr
	for _, pkg := range []string{"systemd-amd64-239", "systemd-amd64-100"} {
		rtmpdir, err := ioutil.TempDir("", "distritest")
		if err != nil {
			return err
		}
		defer os.RemoveAll(rtmpdir)
		meta, err := readMeta(filepath.Join(env.DefaultRepo, systemd+".meta.textproto"))
		if err != nil {
			return err
		}
		// Copy and rename the latest systemd to simulate two versions being present
		for _, dep := range append([]string{systemd}, meta.GetRuntimeDep()...) {
			cp := exec.Command("cp",
				filepath.Join(env.DefaultRepo, dep+".squashfs"),
				filepath.Join(env.DefaultRepo, dep+".meta.textproto"),
				rtmpdir)
			cp.Stderr = os.Stderr
			if err := cp.Run(); err != nil {
				return fmt.Errorf("%v: %v", cp.Args, err)
			}
		}
		idx := strings.LastIndexByte(pkg, '-')
		base, version := pkg[:idx], pkg[idx+1:]
		for _, suffix := range []string{"squashfs", "meta.textproto"} {
			if err := os.Rename(
				filepath.Join(rtmpdir, "systemd-amd64-239."+suffix),
				filepath.Join(rtmpdir, pkg+"."+suffix)); err != nil {
				return err
			}
			if err := os.Symlink(pkg+"."+suffix, filepath.Join(rtmpdir, base+"."+suffix)); err != nil {
				return err
			}
		}

		metaFn := filepath.Join(rtmpdir, pkg+".meta.textproto")
		pm, err := readMeta(metaFn)
		if err != nil {
			return err
		}
		pm.Version = proto.String(version)
		if err := ioutil.WriteFile(metaFn, []byte(proto.MarshalTextString(pm)), 0644); err != nil {
			return err
		}
		addr, cleanup, err := export(ctx, rtmpdir)
		if err != nil {
			return err
		}
		defer cleanup()
		addrs[pkg] = addr
	}

	ctmpdir, err := ioutil.TempDir("", "distritest")
	if err != nil {
		return err
	}
	defer os.RemoveAll(ctmpdir)
	reposd := filepath.Join(ctmpdir, "repos.d")
	if err := os.Mkdir(reposd, 0755); err != nil {
		return err
	}
	for pkg, addr := range addrs {
		if err := ioutil.WriteFile(filepath.Join(reposd, pkg), []byte("http://"+addr), 0644); err != nil {
			return err
		}
	}
	install := exec.Command("distri",
		append([]string{
			"install",
			"-root=" + tmpdir,
		}, pkg...)...)
	install.Env = []string{"DISTRICFG=" + ctmpdir}
	install.Stderr = os.Stderr
	install.Stdout = os.Stdout
	if err := install.Run(); err != nil {
		return fmt.Errorf("%v: %v", install.Args, err)
	}
	return nil
}

func TestInstall(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		desc        string
		installFunc func(ctx context.Context, tmpdir string, pkg ...string) error
		pkgsFull    []string
		pkgs        []string
	}{
		{
			desc:        "File",
			installFunc: installFile,
			pkgsFull:    []string{systemd},
			pkgs:        []string{systemd},
		},

		{
			desc:        "HTTP",
			installFunc: installHTTP,
			pkgsFull:    []string{systemd},
			pkgs:        []string{systemd},
		},

		{
			desc:        "HTTPResolveVersion",
			installFunc: installHTTP,
			pkgsFull:    []string{systemd},
			pkgs:        []string{"systemd-amd64"},
		},

		{
			desc:        "HTTPResolveAll",
			installFunc: installHTTP,
			pkgsFull:    []string{systemd},
			pkgs:        []string{"systemd"},
		},

		{
			desc:        "HTTPMultiplePkgs",
			installFunc: installHTTPMultiple,
			pkgsFull:    []string{systemd, bash},
			pkgs:        []string{systemd, bash},
		},

		{
			desc:        "HTTPMultipleVersions",
			installFunc: installHTTPMultipleVersions,
			pkgsFull:    []string{systemd},
			pkgs:        []string{"systemd"},
		},
	} {
		tt := tt // copy
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			ctx, canc := context.WithCancel(context.Background())
			defer canc()

			// install a package from DISTRIROOT/build/distri/pkg to a temporary directory
			tmpdir, err := ioutil.TempDir("", "integrationinstall")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tmpdir)

			if err := tt.installFunc(ctx, tmpdir, tt.pkgs...); err != nil {
				t.Fatal(err)
			}

			for _, pkg := range tt.pkgsFull {
				m, err := readMeta(filepath.Join(env.DefaultRepo, pkg+".meta.textproto"))
				if err != nil {
					t.Fatal(err)
				}
				for _, pkg := range append([]string{pkg}, m.GetRuntimeDep()...) {
					t.Run("VerifyPackageInstalled/"+pkg, func(t *testing.T) {
						if _, err := os.Stat(filepath.Join(tmpdir, "roimg", pkg+".squashfs")); err != nil {
							t.Fatal(err)
						}
						if _, err := os.Stat(filepath.Join(tmpdir, "roimg", pkg+".meta.textproto")); err != nil {
							t.Fatal(err)
						}
					})
				}
			}

			t.Run("VerifyEtcCopiedSystemd", func(t *testing.T) {
				if _, err := os.Stat(filepath.Join(tmpdir, "etc", "systemd", "system.conf")); err != nil {
					t.Fatal(err)
				}
				linkName := filepath.Join(tmpdir, "etc", "xdg", "systemd", "user")
				st, err := os.Lstat(linkName)
				if err != nil {
					t.Fatal(err)
				}
				if st.Mode()&os.ModeSymlink == 0 {
					t.Errorf("%s unexpectedly not a symbolic link", linkName)
				}
				target, err := os.Readlink(linkName)
				if err != nil {
					t.Fatal(err)
				}
				if got, want := target, "../../systemd/user"; got != want {
					t.Errorf("unexpected link target: got %q, want %q", got, want)
				}
			})

			t.Run("VerifyEtcCopiedGlibc", func(t *testing.T) {
				if _, err := os.Stat(filepath.Join(tmpdir, "etc", "rpc")); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
