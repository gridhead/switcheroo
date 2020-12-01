# switcheroo
An easy-to-use CPU governor switcher - Run things fast or save some juice!

## Usage
1. Download a binary release according to your device's system architecture.
2. Make the binary release executable by running `chmod +x switcheroo`.
3. View the help and support section for script usage by running `./switcheroo -help`.
    ```shell script
    $ ./switcheroo -help
    [ i ] SWTICHEROO - Run things fast or save some juice!
          © 2019-2020 Akashdeep Dhar <t0xic0der@fedoraproject.org>
          -crnt - Know the currently selected CPU governor
          -help - Read about the scriplet and its creator
          -list - List the collection of all available CPU governors
          -setn - Change CPU governor by name - requires superuser
    ```
4. List all available CPU governors for your device by running `./switcheroo -list`.
    ```shell script
    $ ./switcheroo -list 
    [ ✓ ] Available CPU governors were successfully read
          conservative
          ondemand
          userspace
          powersave
          performance
          schedutil
    ```
5. Know about the currently selected CPU governor by running `./switcheroo -crnt`.
    ```shell script
    [alarm@alarm bash]$ ./switcheroo -crnt
    [ ✓ ] performance is the currently selected CPU governor
    ```
6. Switch to another CPU governor of your choice by running `./switcheroo -setn <cpu-governor-name>` with root permissions.
    ```shell script
    $ sudo ./switcheroo -setn powersave
    [ ✓ ] powersave replaces performance as your current CPU governor
    ```
    or
    ```shell script
    # ./switcheroo -setn powersave
    [ ✓ ] powersave replaces performance as your current CPU governor
    ```

## Downloads
1. For generic desktop PCs, [click here](https://github.com/t0xic0der/switcheroo/releases/download/v0.1.0/switcheroo-v0.1.0-amd64).
2. For Raspberry Pi 3B and 4B, [click here](https://github.com/t0xic0der/switcheroo/releases/download/v0.1.0/switcheroo-v0.1.0-aarch64).