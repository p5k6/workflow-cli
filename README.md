|![](https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Anchor_pictogram_yellow.svg/156px-Anchor_pictogram_yellow.svg.png) | Hephy Workflow is the open source fork of Deis Workflow.<br />Please [go here](https://www.teamhephy.com/) for more detail. |
|---:|---|
| 01/27/2026 | Team Hephy appears to have abandoned project; p5k6 fork updated with current libraries, Go 1.21, and modern TLS/security stack |
| 08/27/2018 | Team Hephy [blog][] comes online |
| 08/20/2018 | Deis [#community slack][] goes dark |
| 08/10/2018 | Hephy Workflow [v2.19.4][] fourth patch release |
| 08/08/2018 | [Deis website][] goes dark, then redirects to Azure Kubernetes Service |
| 08/01/2018 | Hephy Workflow [v2.19.3][] third patch release |
| 07/17/2018 | Hephy Workflow [v2.19.2][] second patch release |
| 07/12/2018 | Hephy Workflow [v2.19.1][] first patch release |
| 06/29/2018 | Hephy Workflow [v2.19.0][] first release in the open source fork of Deis |
| 06/16/2018 | Hephy Workflow [v2.19][] series is announced |
| 03/01/2018 | End of Deis Workflow maintenance: critical patches no longer merged |
| 12/11/2017 | Team Hephy [slack community][] invites first volunteers |
| 09/07/2017 | Deis Workflow [v2.18][] final release before entering maintenance mode |
| 09/06/2017 | Team Hephy [slack community][] comes online |

# Deis Client

[![Go Report Card](https://goreportcard.com/badge/github.com/p5k6/workflow-cli)](https://goreportcard.com/report/github.com/p5k6/workflow-cli)
[![codecov](https://codecov.io/gh/deis/workflow-cli/branch/master/graph/badge.svg)](https://codecov.io/gh/deis/workflow-cli)

`deis` is a command line utility used to interact with the [Deis](http://deis.io) open source PaaS.

Please add any [issues](https://github.com/p5k6/workflow-cli/issues) you find with this software to the [Deis Workflow CLI Project](https://github.com/p5k6/workflow-cli).

## Installation

### Prebuilt binaries - *no longer available*

### From Scratch on OS X and Linux

To compile the client from scratch, ensure you have Docker installed and run

    $ make

### From Scratch on Windows

(ed note - I'm on macos and have not tested windows myself....)

To compile the client from scratch, open PowerShell and execute the following commands in the source directory.

    $ .\make bootstrap
    $ .\make build

`.\make bootstrap` will fetch all required dependencies, while `.\make build` will compile and install
the client in the current directory.

    $ .\deis --version

## Usage

Running `deis help` will give you a up to date list of `deis` commands.
To learn more about a command run `deis help <command>`.

## License

see [LICENSE](https://github.com/p5k6/workflow-cli/blob/master/LICENSE)

[v2.18]: https://github.com/teamhephy/workflow/releases/tag/v2.18.0
[k8s-home]: http://kubernetes.io
[install-k8s]: http://kubernetes.io/gettingstarted/
[mkdocs]: http://www.mkdocs.org/
[issues]: https://github.com/teamhephy/workflow/issues
[prs]: https://github.com/teamhephy/workflow/pulls
[Deis website]: http://deis.com/
[blog]: https://blog.teamhephy.info/blog/
[#community slack]: https://slack.deis.io/
[slack community]: https://slack.teamhephy.com/
[v2.18]: https://github.com/teamhephy/workflow/releases/tag/v2.18.0
[v2.19]: https://web.teamhephy.com
[v2.19.0]: https://gist.github.com/Cryptophobia/24c204583b18b9fc74c629fb2b62dfa3/revisions
[v2.19.1]: https://github.com/teamhephy/workflow/releases/tag/v2.19.1
[v2.19.2]: https://github.com/teamhephy/workflow/releases/tag/v2.19.2
[v2.19.3]: https://github.com/teamhephy/workflow/releases/tag/v2.19.3
[v2.19.4]: https://github.com/teamhephy/workflow/releases/tag/v2.19.4
