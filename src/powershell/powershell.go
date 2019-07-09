package powershell

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
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
	out, err := exec.Command("cmd", "powershell -command \""+pscommand+"\"").Output()

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
	out, err := exec.Command("cmd", "powershell -command \""+pscommand+"\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("Command Successfully Executed")

	fmt.Println("Starting install of image")
	pscommand = "Add-AppxPackage .\\" + appName
	out, err = exec.Command("cmd", "powershell -command \""+pscommand+"\"").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is %s\n", out)
	fmt.Println("Install of image ended, starting initialization")
	pscommand = "bash"
	out, err = exec.Command("cmd", "powershell -command \""+pscommand+"\"").Output()

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

	out, err := exec.Command("cmd", "powershell -command \""+strings.Join(pscommand, ";")+"\"").Output()

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

func runWSLPowershellScript(args []string) string {
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	pscommand := make([]string, 1)
	pscommand[0] = "wsl.exe bash -c \""
	pscommand = append(pscommand, args...)
	//pscommand = append(pscommand, "\" 2>&1 | %{\"$_\"}") //to output to console instead of error object
	pscommand = append(pscommand, "\"") //needed for closing

	stdout, stderr, err := shell.Execute(strings.Join(pscommand, ""))
	if err != nil {
		fmt.Println(err)
		panic(err)

	} else {
		//noinspection GoNilness
		return strings.Join([]string{stdout, stderr, err.Error()}, "")
	}
}

type iPerfObject struct {
	address   string //-c client
	port      int
	bandwidth int //target bandwidth, default tcp unlimited, udp 1 Mbps. AKA speed, shouldn't exceed 80% of line speed.

	//flags
	verbose bool
	reverse bool //Server sends, client receives, default: true
	time    int  //seconds to transmit for
	omit    int  //how many seconds to omit for for TCP slow start, default: 2
	isUdp   bool //udp or tcp, default false
}

func NewIPerfObject(address string, port int, bandwidth int, verbose bool, reverse bool, time int, omit int, isUdp bool) iPerfObject {
	return iPerfObject{address: address, port: port, bandwidth: bandwidth, verbose: verbose, reverse: reverse, time: time, omit: omit, isUdp: isUdp}
}

func iPerfTest(object iPerfObject) string {
	iPerfTest := []string{"iperf3"}
	iPerfTest = append(iPerfTest, "-4") //IPv4
	if (iPerfObject{}.address != "") {
		iPerfTest = append(iPerfTest, "--reverse")
	} else {
		log.Println("failure, no address given")
		return "Failed iPerf Test, no address given"
	}
	if (iPerfObject{}.port != 0) { //default port is 5001, but rare
		iPerfTest = append(iPerfTest, " --port "+strconv.Itoa(iPerfObject{}.port))
	} else {
		iPerfTest = append(iPerfTest, " --port 5001 ")
	}
	if (iPerfObject{}.bandwidth != 0) { //default port is 5001, but rare
		iPerfTest = append(iPerfTest, " --bandwidth "+strconv.Itoa(iPerfObject{}.bandwidth))
	}
	if (iPerfObject{}.verbose) {
		iPerfTest = append(iPerfTest, " --verbose ")
	}
	if (iPerfObject{}.reverse) { //reverse = client receives/download test
		iPerfTest = append(iPerfTest, "--reverse")
	}
	if (iPerfObject{}.time != 0) {
		iPerfTest = append(iPerfTest, "--time "+strconv.Itoa(iPerfObject{}.time))
	}
	if (iPerfObject{}.omit != 0) {
		iPerfTest = append(iPerfTest, "--omit "+strconv.Itoa(iPerfObject{}.omit))
	}
	if (iPerfObject{}.isUdp) { //if false, default is tcp, no addition to command for iPerf3
		iPerfTest = append(iPerfTest, "--udp ")
	}

	return runWSLPowershellScript([]string{strings.Join(iPerfTest, " ")}) //note the space in Join separator.
}

func RunIperf3Output(address string, port string, speed string) string {
	return runWSLPowershellScript([]string{"iperf3 -4 -V -R -t 5 -O 3 -u -b " + speed + "M -c " + address + " -p " + port})
}

func RunIPerf3Test(i iPerfObject) {
	iPerfTest(i)
}

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
	stdout, stderr, err := shell.Execute("wsl.exe bash -c \"iperf3 -4 -V -R -t 5 -O 3 -u -b " + speed + "M -c " + address + " -p " + port + "\" 2>&1 | %{\"$_\"}")
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
