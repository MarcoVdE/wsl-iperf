package powershell


import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func main() {
}

func startPowerShellSession() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	//return shell

	// ... and interact with it
	stdout, stderr, err := shell.Execute("Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux")
	if err != nil {
		panic(err)
	}
	fmt.Println(stdout, stderr)
}

func EnableWSL() {
	fmt.Println("Enabling WSL")
	//startPowerShellSession()
	pscommand := "Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux"
	out, err := exec.Command("cmd", "powershell -command \"" + pscommand + "\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("Command Successfully Executed")
}

func InstallUbuntuWSL() {
	appName := "wsl-ubuntu-1804"
	fmt.Println("Installing Ubuntu WSL")
	fmt.Println("Initializing download")
	pscommand := "Invoke-WebRequest -Uri https://aka.ms/" + appName + " -OutFile Ubuntu.appx -UseBasicParsing"
	out, err := exec.Command("cmd", "powershell -command \"" + pscommand  + "\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("Command Successfully Executed")

	fmt.Println("Starting install of image")
	pscommand = "Add-AppxPackage .\\" + appName
	out, err = exec.Command("cmd", "powershell -command \"" + pscommand  + "\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("Install of image ended, starting initialization")
	pscommand = "bash"
	out, err = exec.Command("cmd", "powershell -command \"" + pscommand  + "\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("ended initialization")
}

func InstallIPerf3WSL() {
	fmt.Println("Installing iperf")
	pscommand := make([]string, 1)
	pscommand[0] = "bash"
	pscommand = append(pscommand, "sudo apt-get install iperf3")

	out, err := exec.Command("cmd", "powershell -command \"" + strings.Join(pscommand, ";")  + "\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("ended install")
}

//func RunIPerfTest(address string, speed string) {
//	fmt.Println("Starting iperf test")
//	pscommand := make([]string, 1)
//	pscommand[0] = "bash"
//	pscommand = append(pscommand, "iperf3 -4 -V -R -t 60 -O 3 -u -b " + speed + "M -c " + address + " -p 3334 ")
//
//	cmd := exec.Command("cmd", strings.Join(pscommand, ";"))
//
//	stdout, _ := cmd.StdoutPipe()
//	cmd.Start()
//	oneByte := make([]byte, 100)
//	num := 1
//	for {
//		_, err := stdout.Read(oneByte)
//		if err != nil {
//			fmt.Printf(err.Error())
//			break
//		}
//		r := bufio.NewReader(stdout)
//		line, _, _ := r.ReadLine()
//		fmt.Println(string(line))
//		num = num + 1
//		if num > 3 {
//			os.Exit(0)
//		}
//	}
//
//	cmd.Wait()
//}

func RunIperf3(address string, port string, speed string) {

	//fmt.Print(backend.Local{})
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
	panic(err)
	}
	defer shell.Exit()

	pscommand := make([]string, 1)
	pscommand[0] = "bash"
	//pscommand = append(pscommand, "iperf3 -4 -V -R -t 60 -O 3 -u -b " + speed + "M -c " + address + " -p 3334 ")

	// ... and interact with it
	//stdout, stderr, err := shell.Execute(strings.Join(pscommand, ";"))
	//stdout, stderr, err := shell.Execute("wsl.exe bash -c \"ping google.com" + "\"")
	//if err != nil {
	//	panic(err)
	//}
	stdout, stderr, err := shell.Execute("wsl.exe bash -c \"iperf3 -4 -V -R -t 5 -O 3 -u -b " + speed + "M -c " + address +" -p " + port +"\" 2>&1 | %{\"$_\"}")
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	//stdout, stderr, err = shell.Execute("iperf3 -4 -V -R -t 5 -O 3 -u -b 20M -c iperf.atomic.ac -p 3334")
	//stdout, stderr, err = shell.Execute("ping google.com")
	//if err != nil {
	//	panic(err)
	//}


	fmt.Println(stdout, " stderr:", stderr)
}
