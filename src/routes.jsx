import {
  HomeIcon,
  UserCircleIcon,
  TableCellsIcon,
  InformationCircleIcon,
  ServerStackIcon,
  RectangleStackIcon,
  NoSymbolIcon,
} from "@heroicons/react/24/solid";
import { Home, Profile, Blogs, Notifications, Editor } from "@/pages/dashboard";
import { SignIn, SignUp } from "@/pages/auth";
import NotFound from "@/pages/errors/not-found";

const icon = {
  className: "w-5 h-5 text-inherit",
};

export const routes = [
  {
    layout: "dashboard",
    pages: [
      {
        icon: <HomeIcon {...icon} />,
        name: "dashboard",
        path: "/home",
        element: <Home />,
      },
      {
        icon: <UserCircleIcon {...icon} />,
        name: "profile",
        path: "/profile",
        element: <Profile />,
      },
      {
        icon: <TableCellsIcon {...icon} />,
        name: "Blogs Manager",
        path: "/blogs",
        element: <Blogs />,
      },
      {
        icon: <InformationCircleIcon {...icon} />,
        name: "notifications",
        path: "/notifications",
        element: <Notifications />,
      },
      {
        icon: <NoSymbolIcon {...icon} />,
        name: "editor",
        path: "/editor/:id?",
        element: <Editor />,
      },
      {
        icon: <NoSymbolIcon {...icon} />,
        name: "Not Found",
        path: "/not-found",
        element: <NotFound />,
      },
    ],
  },
  {
    title: "auth pages",
    layout: "auth",
    pages: [
      {
        icon: <ServerStackIcon {...icon} />,
        name: "sign in",
        path: "/sign-in",
        element: <SignIn />,
      },
      {
        icon: <RectangleStackIcon {...icon} />,
        name: "sign up",
        path: "/sign-up",
        element: <SignUp />,
      },
    ].filter((item) => !["/sign-up"].includes(item.path)),
  },
  {
    title: "errors pages",
    layout: "errors",
    pages: [
      {
        icon: <NoSymbolIcon {...icon} />,
        name: "Not Found",
        path: "/not-found",
        element: <NotFound />,
      },
    ],
  },
];

export default routes;
