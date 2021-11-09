package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	pbmatcher "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	iproto "github.com/lyft/xdsmatcher/internal/proto"
	"github.com/lyft/xdsmatcher/pkg/matcher"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	matcherYamlFile string
	inputYamlFiles  []string
)

func init() {
	evalCommand.PersistentFlags().StringVar(&matcherYamlFile, "matcher-yaml", "", "path to file containing a YAML describing the Matcher proto to use")
	evalCommand.PersistentFlags().StringArrayVar(&inputYamlFiles, "input-yamls", []string{}, "list of input files to validate")

	evalCommand.MarkFlagRequired("matcher-yaml")
	evalCommand.MarkFlagRequired("input-yamls")

	rootCmd.AddCommand(evalCommand)
}

var evalCommand = &cobra.Command{
	Use:   "eval",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		m := &pbmatcher.Matcher{}
		err := protoFromYamlFile(matcherYamlFile, m)
		if err != nil {
			panic(err)
		}

		matcher, err := matcher.Create(m)
		if err != nil {
			panic(err)
		}

		for _, f := range inputYamlFiles {
			a := &anypb.Any{}
			err = protoFromYamlFile(f, a)
			if err != nil {
				panic(err)
			}

			m, err := anypb.UnmarshalNew(a, proto.UnmarshalOptions{})
			if err != nil {
				panic(err)
			}
			r, err := matcher.Match(m)
			if err != nil {
				panic(err)
			}
			switch {
			case r.NeedMoreData:
				fmt.Printf("unable to match, insufficient data\n")
			case r.MatchResult == nil:
				fmt.Printf("result: no match\n")
			case r.MatchResult.Action == nil:
				fmt.Printf("result: no match\n")
			default:
				fmt.Printf("result: %v\n", r.MatchResult.Action)
			}
		}

	},
}

func protoFromYamlFile(path string, m proto.Message) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = iproto.ProtoFromYaml(d, m)
	if err != nil {
		return err
	}

	return nil
}
