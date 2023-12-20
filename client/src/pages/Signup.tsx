import { useContext, useState } from "react";
import { toast } from "react-hot-toast";
import { Link, useNavigate } from "react-router-dom";
import { UserContext } from "../contexts/UserContext";
import { useAuth } from "../hooks/useAuth";
import { AuthService } from "../services/AuthService";

const Signup = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();
  const { login } = useAuth();
  const { setUserDetails } = useContext(UserContext);

  const handleSignup = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const { token, user } = await AuthService.signup(name, email, password);
      login(token);
      setUserDetails(user.id, user.name, user.email);
      toast.success("Signed up successfully.");
      navigate("/");
    } catch (error) {
      if (error instanceof Error) {
        toast.error(`Sign up failed. ${error.message}`);
      } else {
        toast.error("Sign up failed.");
      }
      console.error("Error:", error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="p-6 bg-white shadow-md rounded w-full max-w-md mx-auto">
        <form onSubmit={handleSignup}>
          <h1 className="text-2xl font-semibold mb-4">
            Sign Up for an Account
          </h1>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Name"
            className="w-full p-2 mb-3 border rounded"
          />
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
            disabled={!name || !email || !password}
          >
            Sign Up
          </button>
          <p className="mt-4">
            Already have an account?
            <Link
              to="/login"
              className="text-blue-500 ml-1 hover:text-blue-600"
            >
              Login
            </Link>
          </p>
        </form>
      </div>
    </div>
  );
};

export default Signup;
