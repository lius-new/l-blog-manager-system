import { login } from "@/libs/action";
import {
  Input,
  Button,
  Typography,
  Dialog,
  DialogHeader,
  DialogBody,
  DialogFooter,
} from "@material-tailwind/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export function SignIn() {
  const navigate = useNavigate();
  const [errOpen, setErrOpen] = useState(false);
  const [loginForm, setLoginForm] = useState({
    username: "",
    password: "",
  });

  const loginHandle = async () => {
    try {
      const res = await login(loginForm.username, loginForm.password);
      if (res.status) navigate("/dashboard/home");
      else errOpenHandle();
    } catch (err) {
      errOpenHandle();
    }
  };

  const errOpenHandle = () => {
    document.cookie = "secret=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    setErrOpen(!errOpen);
  };

  return (
    <>
      <section className="flex gap-4">
        <div className="w-full lg:w-3/5 flex flex-col items-center justify-center">
          <div className="text-center">
            <Typography variant="h2" className="font-bold mb-4">
              LIUS BACEND PAGE
            </Typography>
          </div>
          <form className="mt-8 mb-2 mx-auto w-80 max-w-screen-lg lg:w-1/2">
            <div className="mb-1 flex flex-col gap-6">
              <Typography
                variant="small"
                color="blue-gray"
                className="-mb-3 font-medium"
              >
                邮箱
              </Typography>
              <Input
                size="lg"
                placeholder="name@mail.com"
                className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                labelProps={{
                  className: "before:content-none after:content-none",
                }}
                onChange={(e) =>
                  setLoginForm({ ...loginForm, username: e.target.value })
                }
              />
              <Typography
                variant="small"
                color="blue-gray"
                className="-mb-3 font-medium"
              >
                密码
              </Typography>
              <Input
                type="password"
                size="lg"
                placeholder="********"
                className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                labelProps={{
                  className: "before:content-none after:content-none",
                }}
                onChange={(e) =>
                  setLoginForm({ ...loginForm, password: e.target.value })
                }
              />
            </div>
            <Button
              className="mt-6"
              fullWidth
              disabled={!loginForm.username || !loginForm.password}
              onClick={loginHandle}
            >
              登陆
            </Button>
            <div className="py-12 text-gray-900 text-right">
              别看了, 这只是我的个人后台页面:) & 小心 BLOCKED...
            </div>
          </form>
        </div>
        <div className="w-2/5 h-screen hidden lg:block p-8">
          <img
            src="/img/pattern.png"
            className="h-full w-full object-cover rounded-3xl "
          />
        </div>
      </section>
      <Dialog open={errOpen} size="xs" handler={errOpenHandle}>
        <DialogHeader>登录失败</DialogHeader>
        <DialogBody>因为某些原因登录失败, 请联系管理查看服务器日志</DialogBody>
        <DialogFooter>
          <Button variant="gradient" color="green" onClick={errOpenHandle}>
            <span>知晓</span>
          </Button>
        </DialogFooter>
      </Dialog>
    </>
  );
}

export default SignIn;
