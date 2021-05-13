"""Role testing files using testinfra."""


def test_hosts_file(host):
    """Validate /etc/hosts file."""
    f = host.file("/etc/hosts")

    assert f.exists
    assert f.user == "root"
    assert f.group == "root"


def test_app_user_ispresent(host):
    u = host.user("speedtest")

    print(u)
    assert u.exists
    assert u.expiration_date is None


def test_app_service_file_ispresent(host):
    f = host.file("/etc/systemd/system/speedtest.service")

    assert f.exists
    assert f.user == "root"
    assert f.group == "root"

    assert f.contains(
        "ExecStart="
        "/app/speedtest_wrapper/speedtest-wrapper-go "
        "test "
        "-p "
        "--config /app/speedtest_wrapper/config.yml"
    )

    assert f.contains("User=speedtest")


def test_app_timer_file_ispresent(host):
    f = host.file("/etc/systemd/system/speedtest.timer")

    assert f.exists
    assert f.user == "root"
    assert f.group == "root"

    assert f.contains("OnCalendar=\\*\\-\\*\\-\\* \\*:\\*:00")
