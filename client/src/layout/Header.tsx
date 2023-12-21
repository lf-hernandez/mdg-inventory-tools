import { useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";

import mdgLogo from "../assets/logo.webp";
import { UserAvatar } from "../components/UserAvatar";
import { useAuth } from "../hooks/useAuth";
import { useCurrentUser } from "../hooks/useCurrentUser";

export const Header = () => {
  const { user, clearUserDetails } = useCurrentUser();
  const navigate = useNavigate();
  const { logout } = useAuth();

  const [isDropdownVisible, setDropdownVisible] = useState(false);
  const dropdownRef = useRef<HTMLDivElement | null>(null);
  const userAvatarRef = useRef<HTMLButtonElement | null>(null);

  const handleLogout = () => {
    setDropdownVisible(false);
    logout();
    clearUserDetails();
    toast.success("Logged out successfully.");
    navigate("/login");
  };

  const toggleDropdown = () => {
    setDropdownVisible(!isDropdownVisible);
  };

  const closeDropdown = () => {
    setDropdownVisible(false);
  };

  useEffect(() => {
    const handleClickOutside = (event: Event) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node) &&
        userAvatarRef.current &&
        !userAvatarRef.current.contains(event.target as Node)
      ) {
        closeDropdown();
      }
    };

    if (isDropdownVisible) {
      document.addEventListener("mousedown", handleClickOutside);
    } else {
      document.removeEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [isDropdownVisible]);

  return (
    <header className="bg-gray-100 py-4 px-6 flex justify-between items-center">
      <div className="flex flex-row">
        <img src={mdgLogo} alt="MDG Logo" className="h-20 mr-4" />
        <div className="flex flex-col justify-center h-20">
          <h1 className="text-4xl lg:text-5xl">Inventory Manager</h1>
        </div>
      </div>
      {user.name && (
        <div className="flex items-center relative">
          <button
            onClick={toggleDropdown}
            className="p-2 rounded-full hover:bg-gray-300 focus:outline-none"
            ref={userAvatarRef}
          >
            <UserAvatar fullName={user.name} />
          </button>
          <div
            ref={dropdownRef}
            className={`${
              isDropdownVisible ? "block" : "hidden"
            } absolute right-0 mt-2 bg-white border border-gray-300 shadow-lg rounded-lg text-gray-800`}
            style={{ top: "calc(100%)", right: 0 }}
          >
            <div className="p-4">
              <div className="text-lg font-semibold">{user.name}</div>
              <div className="text-sm">{user.email}</div>
            </div>
            <hr className="border-gray-300" />
            <button
              onClick={handleLogout}
              className="block w-full py-2 px-4 text-left hover:bg-gray-100 text-red-500 hover:text-mdg-red font-semibold"
            >
              Logout
            </button>
          </div>
        </div>
      )}
    </header>
  );
};
