import { useNavigate } from "react-router-dom";
import { AddItemForm } from "../components/AddItemForm";
import { ItemList } from "../components/ItemList";
import { SearchForm } from "../components/SearchForm";
import { useAuth } from "../hooks/useAuth";

const Home = () => {
  const navigate = useNavigate();
  const { logout } = useAuth();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <div className="mx-auto max-w-7xl p-4">
      <div className="flex flex-row justify-end">
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
