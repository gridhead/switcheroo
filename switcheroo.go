package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func FR_RED(textobjc string) string {
	return "\u001b[31m" + textobjc + "\u001b[0m"
}

func FR_GREEN(textobjc string) string {
	return "\u001b[32m" + textobjc + "\u001b[0m"
}

func FR_YELLOW(textobjc string) string {
	return "\u001b[33m" + textobjc + "\u001b[0m"
}

func FR_BLUE(textobjc string) string {
	return "\u001b[34m" + textobjc + "\u001b[0m"
}

func GN_TEXT(textobjc string) string {
	return "      " + textobjc
}

func view_currently_selected_governor() {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor")
	otpt, err := cmd.Output()
	if err != nil {
		fmt.Printf(FR_RED("[ ! ]") + " " + "Currently selected CPU governor could not be read")
	} else {
		var curtgvnr string = string(otpt)[0 : len(otpt)-1]
		fmt.Printf(FR_GREEN("[ ✓ ]") + " " + FR_YELLOW(curtgvnr) + " " + "is the currently selected CPU governor" + "\n")
	}
}

func list_available_governors() {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors")
	otpt, err := cmd.Output()
	if err != nil {
		fmt.Printf(FR_RED("[ ! ]") + " " + "Available CPU governors could not be read" + "\n")
	} else {
		var gvnrlist []string = strings.Fields(string(otpt))
		fmt.Printf(FR_GREEN("[ ✓ ]"+" "+"Available CPU governors were successfully read") + "\n")
		for indx := 0; indx < len(gvnrlist); indx++ {
			fmt.Println(GN_TEXT(FR_YELLOW(gvnrlist[indx])))
		}
	}
}

func set_governor_by_name(gvnrname string) {
	var gosg bool = false
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors")
	otpt, err := cmd.Output()
	if err != nil {
		fmt.Printf(FR_RED("[ ! ]") + " " + "Available governors could not be read" + "\n")
	} else {
		var gvnrlist []string = strings.Fields(string(otpt))
		for indx := 0; indx < len(gvnrlist); indx++ {
			if gvnrname == gvnrlist[indx] {
				gosg = true
				break
			}
		}
	}
	if gosg == true {
		cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor")
		otpt, err := cmd.Output()
		if err != nil {
			fmt.Printf(FR_RED("[ ! ]") + " " + "Could not compare provided CPU governor name with the existing" + "\n")
		} else {
			var crntgvnr string = string(otpt)[0 : len(otpt)-1]
			if gvnrname == crntgvnr {
				fmt.Printf(FR_RED("[ ! ]") + " " + FR_YELLOW(gvnrname) + " " + "is already your current CPU governor" + "\n")
			} else {
				cmd := exec.Command("bash", "-c", "echo "+gvnrname+" | tee /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor")
				otpt, err := cmd.Output()
				if err != nil {
					fmt.Printf(FR_RED("[ ! ]") + " " + "Provided CPU governor could not be selected" + "\n")
				} else {
					if string(otpt)[0:len(otpt)-1] == gvnrname {
						fmt.Printf(FR_GREEN("[ ✓ ]") + " " + FR_YELLOW(gvnrname) + " " + "replaces" + " " + FR_YELLOW(crntgvnr) + " " + "as your current CPU governor" + "\n")
					} else {
						fmt.Printf(FR_RED("[ ! ]") + " " + "Outcome of CPU governor change could not be verified" + "\n")
					}
				}
			}
		}
	} else {
		fmt.Printf(FR_RED("[ ! ]") + " " + "Provided CPU governor name could not be verified" + "\n")
	}
}

func help_and_support() {
	fmt.Printf(FR_BLUE("[ i ] ") + FR_GREEN("SWTICHEROO") + " - Run things fast or save some juice!" + "\n" +
		GN_TEXT("© 2019-2020 Akashdeep Dhar <t0xic0der@fedoraproject.org>") + "\n" +
		GN_TEXT(FR_YELLOW("-crnt")+" - "+"Know the currently selected CPU governor") + "\n" +
		GN_TEXT(FR_YELLOW("-help")+" - "+"Read about the scriplet and its creator") + "\n" +
		GN_TEXT(FR_YELLOW("-list")+" - "+"List the collection of all available CPU governors") + "\n" +
		GN_TEXT(FR_YELLOW("-setn")+" - "+"Change CPU governor by name - requires superuser") + "\n")
}

func main() {
	var list, crnt, help bool
	var setn string
	flag.BoolVar(&crnt, "crnt", false, "Know the currently selected CPU governor")
	flag.BoolVar(&list, "list", false, "List the collection of all available CPU governors")
	flag.StringVar(&setn, "setn", "", "Change CPU governor by name - requires superuser")
	flag.BoolVar(&help, "help", false, "Read about the scriptlet and its creator")
	flag.Parse()
	if list {
		list_available_governors()
	} else if crnt {
		view_currently_selected_governor()
	} else if help {
		help_and_support()
	} else if setn != "" {
		set_governor_by_name(setn)
	} else {
		help_and_support()
	}
}
