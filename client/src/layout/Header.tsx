import { useNavigate } from "react-router-dom";
import mdgLogo from "../assets/logo.webp";
import { UserDropdown } from "../components/UserDropdown";

export const Header = () => {
  const navigate = useNavigate();

  return (
    <header className="bg-gray-100 py-4 px-6 flex justify-between items-center">
      <button
        onClick={() => navigate("/")}
        aria-label="Home"
        className="flex items-center focus:outline-none"
      >
        <img src={mdgLogo} alt="MDG Logo" className="h-20 mr-4" />
        <h1 className="text-4xl lg:text-5xl">Inventory Manager</h1>
      </button>
      <UserDropdown />
    </header>
  );
};
