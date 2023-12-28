import { useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { useCurrentUser } from "../hooks/useCurrentUser";
import { UserAvatar } from "./UserAvatar";

export const UserDropdown = () => {
  const { user, clearUserDetails } = useCurrentUser();
  const navigate = useNavigate();
  const { logout } = useAuth();
  const [isDropdownVisible, setDropdownVisible] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const userAvatarRef = useRef<HTMLButtonElement | null>(null);

  const handleLogout = () => {
    toggleDropdown();
    logout();
    clearUserDetails();
    toast.success("Logged out successfully.");
    navigate("/login");
  };

  const handleAccountSettings = () => {
    navigate("/account-settings");
    toggleDropdown();
  };

  const toggleDropdown = () => {
    setDropdownVisible(!isDropdownVisible);
  };

  useEffect(() => {
    if (!user || !user.name) {
      toast.error("Invalid session. Please login again.");
      logout();
      clearUserDetails();
      navigate("/login");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [user]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node) &&
        userAvatarRef.current &&
        !userAvatarRef.current.contains(event.target as Node)
      ) {
        setDropdownVisible(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <div className="flex items-center relative">
      <button
        onClick={toggleDropdown}
        className="p-2 rounded-full hover:bg-gray-300 focus:outline-none"
        aria-haspopup="true"
        aria-expanded={isDropdownVisible}
        ref={userAvatarRef}
      >
        <UserAvatar fullName={user.name ?? ""} />
      </button>
      <div
        ref={dropdownRef}
        className={`${
          isDropdownVisible ? "block" : "hidden"
        } absolute right-0 mt-2 bg-white border border-gray-300 shadow-lg rounded-lg text-gray-800`}
        style={{ top: "100%" }}
      >
        <div className="p-4">
          <div className="text-lg font-semibold">{user.name}</div>
          <div className="text-sm">{user.email}</div>
        </div>
        <hr className="border-gray-300" />
        <button
          onClick={handleAccountSettings}
          className="block w-full py-2 px-4 text-left text-black-500 hover:bg-gray-100 hover:text-black-600 font-semibold"
        >
          Account settings
        </button>
        <hr className="border-gray-300" />
        <button
          onClick={handleLogout}
          className="block w-full py-2 px-4 text-left text-red-500 hover:bg-gray-100 hover:text-red-600 font-semibold"
        >
          Logout
        </button>
      </div>
    </div>
  );
};
