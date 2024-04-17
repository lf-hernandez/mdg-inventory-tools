import { Toaster } from "react-hot-toast";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import { AuthProvider } from "./contexts/AuthContext";
import { UserProvider } from "./contexts/UserContext";
import "./index.css";
import { Layout } from "./layout/Layout";
import AccountSettings from "./pages/AccountSettings";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import { ProtectedRoute } from "./router/ProtectedRoute";
import { AnalyticsProvider } from "./contexts/AnalyticsContext";
import Error from "./pages/Error";

const router = createBrowserRouter([
  {
    id: "root",
    path: "/",
    element: <Layout />,
    errorElement: <Error />,
    children: [
      {
        index: true,
        element: <ProtectedRoute component={Home} />,
      },
      {
        path: "account-settings",
        element: <ProtectedRoute component={AccountSettings} />,
      },
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "signup",
        element: <Signup />,
      },
    ],
  },
]);

function App() {
  return (
    <AuthProvider>
      <UserProvider>
        <AnalyticsProvider>
          <RouterProvider router={router} />
          <Toaster />
        </AnalyticsProvider>
      </UserProvider>
    </AuthProvider>
  );
}

export default App;
