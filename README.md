# ordertujuan
Pesanan teman kuliah Ali

install
$ go get -u github.com/labstack/echo

untuk service agar jalan di background go lang
copy file angkot.service ke /lib/systemd/system
chmod 755 ordertujuan.service
systemctl enable ordertujuan.service
systemctl start ordertujuan


cd /etc/init.d
chmod 755 ordertujuanrun
update-rc.d ordertujuanrun defaults
