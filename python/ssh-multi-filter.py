# SSH to Multiple Devices from devices file
import warnings
from cryptography.utils import CryptographyDeprecationWarning
with warnings.catch_warnings():
    warnings.filterwarnings('ignore', category=CryptographyDeprecationWarning)
    import paramiko

from netmiko import ConnectHandler

import sys
import getpass

user = input("Enter your username: ")
password = getpass.getpass("Enter your password: ") #Hide password

path = 'sshMultiRouterTest3.txt' #Text file to dump ouput
sys.stdout = open(path, 'w')

with open('All-Switch-IPs.txt') as routers: #text file in same directory with IP adresses seperated by line (enter key)
    for IP in routers:
        Router = {
            'device_type': 'cisco_ios',
            'ip': IP,
            'username': user,
            'password': password
        }

        net_connect = ConnectHandler(**Router)
        
        switch4500 = net_connect.send_command('sh idprom chass | inc Base')
        switch9400 = net_connect.send_command('sh idprom sup | inc Base')
        switch3850 = net_connect.send_command('sh switch | inc Switch/Stack')

        check1 = switch4500.rfind('M')
        check2 = switch9400.rfind('M')
        check3 = switch3850.rfind('S')
        
        if check1 == 1: #check if chass is valid first
            print(switch4500)
        elif check2 == 1: #fallback to sup
            print(switch9400)
        elif check3 == 7: #fallback to switch
            print(switch3850) 
        else: #something unexpected happened
            print('error'+str(check3))

# Finally close the connection
net_connect.disconnect()