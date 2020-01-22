from icube-cybervision-plugin import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting icube-cybervision-plugin normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        cvbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("icube-cybervision-plugin is running"))
        exit_code = cvbeat_proc.kill_and_wait()
        assert exit_code == 0
