## IDS USING SNORT
Cloud IDS is an intrusion detection service that provides threat detection for intrusions, malware, spyware, and command-and-control attacks on your network. Cloud IDS works by creating a Google-managed peered network with mirrored virtual machine (VM) instances.
I have done the following steps myself and implemented it.
WINDOWS:

## Installation:
1] Download winpcap, npcap - packet sniffing libraries for windows.Both versions are   needed together to run snort.

2]Install snort.exe and snort.rules in registered option.Version for both should be same in        order to use snort.

3]Snort rules unzips etc,preproc_rules,rules,so_rules.Copy rules and preproc_rles from snort rules to snort folder as snort folder has empt or old version rules.

## Edit the snort.conf file:
1]Go to snort-etc-snort.conf to edit. 
2]Set the home net as 192.168.0.0/24 ie network id.(line 45-3 common last change)
3]Set external net as everything except home net.(line 48)
4]Set line 104,105,106 for rules,so_rules and preproc_rules as absolute path for windows.
5]Set line 113,114 for whitelist and blacklist as absolute path for windows.
6]Edit line 186 to include snort log.
7]Edit path(backslash) for line 247 ie snort-lib-snort dynamicpreprocessor(to process the incoming data and send to detection engine).
8]Edit path(backslash) for line 250 ie snort-lib-snort dynamicengine(to check packets with rules specified to detect kya hai).
9]Comment 265,266,267,268-ipv4,ipv6 lines not needed.
10]Comment 335 not needed.
11]Remove comment from 418-we need portscan detection.
12]Go to snort-rules-blacklist.rules it is present.Now edit it to whitelist.rules and save as whitelist.rules as whitelist is not there initially.
13]Rename whitelist , blacklist name properly in line 511,512.
14]For step 7 and step 8 in snort.conf file change all forward slash to backward slash as windows need backward slash.(line 546 - 651 , line 659-661)
15]Remove comment from 659,660,661 as preproc_rules are required.

## Testing snort.
1]Go to cmd prompt run as administrator go to c-snort-bin, run two commands snort -V to check version and snort -W to check available interface on PC.
2] Write this command snort -i 1 -c c:\Snort\etc\snort.conf -T where -i is interface 1 is interface 1 -T is testing -c is open file given.It will check properly all rules and atlast snort successfully validated the configuration will appear.
3]Go to snort-rules-local.rules and paste this:
alert icmp any any -> any any (msg:"Testing ICMP"; sid:1000001;)
alert tcp any any -> any any (msg:"Testing TCP"; sid:1000002;)
alert udp any any -> any any (msg:"Testing udp"; sid:1000003;)

Rule action, protocol ,source ip address, source port, direction operator, destination ip address, destination port, rule options(msg,sid-signature id,rev-version)
alert - generate an alert using the selected alert method, and then log the packet
log - log the packet
pass - ignore the packet
activate - alert and then turn on another dynamic rule
dynamic - remain idle until activated by an activate rule, then act as a log rule

4]Write this command if u want to see all alerts in a file:
snort -i 4 -c c:\Snort\etc\snort.conf -A console > c:\Snort\log\pingtest.txt where -A means run command and print it in file given.
5]Write this command if u want to see all alerts on command prompt:
snort -i 4 -c c:\Snort\etc\snort.conf -A console where-A means displaying op and console means command prompt.
6]Open youtube in chrome to get UDP and boa.com to get TCP.

