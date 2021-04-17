#!/usr/bin/env python3
import unittest
import docker
import os
import time
import subprocess
client = docker.from_env()

class TestDocktor(unittest.TestCase):
    def test_healthcheck(self):
        os.system('.././bin/./docktor healthcheck')
        # waits for container to start then show health
        time.sleep(15)
        containers = client.containers.list()
        for container in containers:
            if 'Health' in container.attrs['State']:
                out = container.attrs['State']['Health']['Status']
                self.assertEqual(out, 'healthy')
    
    def test_heal(self):
        containers = client.containers.list()
        for c in containers:
            container = c.attrs['State']
            if 'Health' in container and container['Health']['Status'] != 'healthy':
                res = subprocess.check_output('.././bin/./docktor heal ' + c.id, shell=True)
                #print('RES', res)
            else:
                print(c.name + ' is not unhealthy')
    
    def test_aheal(self):
        containers = client.containers.list()
        for c in containers:
            os.system('.././bin/./docktor autoheal ' + c.name)
            #res = subprocess.check_output('.././bin/./docktor autoheal ' + c.name, shell=True)
            print('autohealed ' + c.name)

    def test_scan(self):
        containers = client.containers.list()
        for c in containers:
            os.system('.././bin/./docktor scan ' + c.id)

if __name__ == "__main__":
    unittest.main()