package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/patrick/internal/pkg/entity/project"
	"github.com/blackhorseya/patrick/internal/pkg/infra/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init [name]",
	Aliases: []string{"create"},
	Short:   "Initialize a Project",
	Long:    `Initialize a Project`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := log.New(level, output)
		cobra.CheckErr(err)

		ctx := contextx.BackgroundWithLogger(logger)

		projectPath, err := initializeProject(ctx, args)
		if err != nil {
			ctx.Error(err.Error())
			os.Exit(1)
		}

		ctx.Info("Your Cobra application is ready at", zap.String("path", projectPath))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initializeProject(ctx contextx.Contextx, args []string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	modName, err := getModImportPath()
	if err != nil {
		return "", err
	}

	prj := &project.Info{
		AbsolutePath: wd,
		PkgName:      modName,
		AppName:      path.Base(modName),
	}

	err = projectBiz.InitProject(ctx, prj)
	if err != nil {
		return "", err
	}

	return prj.AbsolutePath, nil
}

func getModImportPath() (string, error) {
	mod, cd, err := parseModInfo()
	if err != nil {
		return "", err
	}

	return path.Join(mod.Path, fileToURL(strings.TrimPrefix(cd.Dir, mod.Dir))), nil
}

func fileToURL(in string) string {
	i := strings.Split(in, string(filepath.Separator))
	return path.Join(i...)
}

func parseModInfo() (Mod, CurDir, error) {
	var mod Mod
	var dir CurDir

	m, err := modInfoJSON("-m")
	if err != nil {
		return Mod{}, CurDir{}, err
	}
	err = json.Unmarshal(m, &mod)
	if err != nil {
		return Mod{}, CurDir{}, err
	}

	// Unsure why, but if no module is present Path is set to this string.
	if mod.Path == "command-line-arguments" {
		return Mod{}, CurDir{}, errors.New("please run `go mod init <MODNAME>` before `patrick init`")
	}

	e, err := modInfoJSON("-e")
	if err != nil {
		return Mod{}, CurDir{}, err
	}
	err = json.Unmarshal(e, &dir)
	if err != nil {
		return Mod{}, CurDir{}, err
	}

	return mod, dir, nil
}

type Mod struct {
	Path, Dir, GoMod string
}

type CurDir struct {
	Dir string
}

func goGet(mod string) error {
	return exec.Command("go", "get", mod).Run()
}

func modInfoJSON(args ...string) ([]byte, error) {
	err := os.Setenv("GO111MODULE", "on")
	if err != nil {
		return nil, err
	}

	cmdArgs := append([]string{"list", "-json"}, args...)
	out, err := exec.Command("go", cmdArgs...).Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}
