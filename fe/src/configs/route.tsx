import MainLayout from "@/components/layout/MainLayout";
import Home from "@/views/Home";
import { createBrowserRouter, redirect } from "react-router-dom";

export const router = createBrowserRouter([
  {
    element: <MainLayout />,
    children: [
      {
        path: "/",
        element: <Home />,
      },
      {
        path: "*",
        loader: () => redirect("/"),
      },
    ],
  },
]);
