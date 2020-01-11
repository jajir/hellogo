#!bin/sh
#
# Deploy binaries to EV3 brick.
#

#!/usr/bin/expect
        spawn scp  /usr/bin/file.txt root@<ServerLocation>:/home
        set pass "Your_Password"
        expect {
        password: {send "$pass\r"; exp_continue}
                  }
