import { useState } from "react";
import { toast } from "react-hot-toast";
import { Link, useNavigate } from "react-router-dom";

import { useAuth } from "../hooks/useAuth";
import { useCurrentUser } from "../hooks/useCurrentUser";
import { AuthService } from "../services/AuthService";
import { useAnalytics } from "../hooks/useAnalytics";

const LoginComponent = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();
  const { onLogin } = useAuth();
  const { setUserDetails } = useCurrentUser();
  const { trackEvent } = useAnalytics();

  const handleLogin = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const { token, user } = await AuthService.login(email, password);
      onLogin(token);
      setUserDetails(user.id, user.name, user.email, user.role);
      toast.success("Logged in successfully");
      navigate("/");
      trackEvent("Login", { success: true });
    } catch (error) {
      toast.error("Login error");
      console.error("Error:", error);
      trackEvent("Login", { success: false });
    }
  };

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="p-6 bg-white shadow-md rounded w-full max-w-md mx-auto">
        <form onSubmit={handleLogin}>
          <h1 className="text-2xl font-semibold mb-4">Login to your account</h1>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            className="w-full p-2 mb-3 border rounded"
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            className="w-full p-2 mb-3 border rounded"
          />
          <button
            type="submit"
            className="w-full p-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            disabled={!email || !password}
          >
            Login
          </button>
          <p className="mt-4">
            Don't have an account?
            <Link
              to="/signup"
              className="text-blue-500 ml-1 hover:text-blue-600"
            >
              Sign up
            </Link>
          </p>
        </form>
      </div>
    </div>
  );
};

export default LoginComponent;
