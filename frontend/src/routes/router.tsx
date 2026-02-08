import { createBrowserRouter } from "react-router-dom";
import Layout from "../components/Layout";
import Login from "../pages/Login";
import Register from "../pages/Register";
import Dashboard from "../pages/Dashboard";
import Questions from "../pages/Questions";
import Quiz from "../pages/Quiz";
import Progress from "../pages/Progress";
import { RequireAuth } from "../auth/RequireAuth";

export const router = createBrowserRouter([
  {
    element: <Layout />,
    children: [
      { path: "/login", element: <Login /> },
      { path: "/register", element: <Register /> },
      {
        path: "/",
        element: (
          <RequireAuth>
            <Dashboard />
          </RequireAuth>
        ),
      },
      {
        path: "/questions",
        element: (
          <RequireAuth>
            <Questions />
          </RequireAuth>
        ),
      },
      {
        path: "/quiz",
        element: (
          <RequireAuth>
            <Quiz />
          </RequireAuth>
        ),
      },
      {
        path: "/progress",
        element: (
          <RequireAuth>
            <Progress />
          </RequireAuth>
        ),
      },
    ],
  },
]);
