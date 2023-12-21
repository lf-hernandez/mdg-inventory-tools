import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { AddItemForm } from "../components/AddItemForm";
import { ItemList } from "../components/ItemList";
import { SearchForm } from "../components/SearchForm";
import { useAuth } from "../hooks/useAuth";
import { useCurrentUser } from "../hooks/useCurrentUser";

const Home = () => {
  const navigate = useNavigate();
  const { logout } = useAuth();
  const { user } = useCurrentUser();

  const handleLogout = () => {
    logout();
    toast.success("Logged out successfully.");
    navigate("/login");
  };

  return (
    <div className="mx-auto max-w-7xl p-4">
      <div className={user.name ? "flex justify-between" : "flex justify-end"}>
        {user.name && (
          <div className="flex items-center">
            <p className="text-lg lg:text-xl">Welcome, {user.name}</p>
          </div>
        )}
        <button
          onClick={handleLogout}
          className="mt-4 p-2 bg-red-500 text-white rounded hover:bg-red-600"
        >
          Logout
        </button>
      </div>
      <SearchForm />
      <br />
      <AddItemForm />
      <br />
      <ItemList />
    </div>
  );
};

export default Home;
