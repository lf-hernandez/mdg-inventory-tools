import { UserDropdown } from "../components/UserDropdown";
import { useCurrentUser } from "../hooks/useCurrentUser";

export const Header = () => {
  const { user } = useCurrentUser();

  return (
    <header className="bg-gray-00 py-4 px-6 flex justify-end">
      {user && user.name ? <UserDropdown /> : null}
    </header>
  );
};
