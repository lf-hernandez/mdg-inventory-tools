import { NavLink, useNavigate } from "react-router-dom";
import mdgLogo from "../assets/logo.webp";
import { UserDropdown } from "../components/UserDropdown";

export const Sidebar = () => {
  const navigate = useNavigate();

  const activeLinkStyle = "bg-blue-500 text-white";

  return (
    <aside className="w-full md:w-64 bg-gray-100 md:h-screen flex flex-col">
      <div className="flex-grow overflow-auto">
        <div className="flex flex-col items-center justify-center p-4 border-b border-gray-700">
          <img
            src={mdgLogo}
            alt="MDG Logo"
            className="h-20 mb-2 cursor-pointer"
            onClick={() => navigate("/")}
          />
          <h1 className="text-xl font-semibold">Inventory Manager</h1>
        </div>
        <nav className="mt-4">
          <NavLink
            to="/"
            className={({ isActive }) =>
              isActive
                ? `block py-2 px-4 ${activeLinkStyle}`
                : "block py-2 px-4"
            }
            end
          >
            Search Inventory
          </NavLink>
          <NavLink
            to="/add-inventory"
            className={({ isActive }) =>
              isActive
                ? `block py-2 px-4 ${activeLinkStyle}`
                : "block py-2 px-4"
            }
          >
            Add Inventory
          </NavLink>
          <NavLink
            to="/view-inventory"
            className={({ isActive }) =>
              isActive
                ? `block py-2 px-4 ${activeLinkStyle}`
                : "block py-2 px-4"
            }
          >
            View Inventory
          </NavLink>
        </nav>
      </div>

      <div className="p-4">
        <UserDropdown />
      </div>
    </aside>
  );
};
