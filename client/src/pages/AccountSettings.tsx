import { useState } from "react";
import { toast } from "react-hot-toast";
import { useCurrentUser } from "../hooks/useCurrentUser";
import { AuthService } from "../services/AuthService";

const AccountSettings = () => {
  const { user } = useCurrentUser();

  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmNewPassword, setConfirmNewPassword] = useState("");

  const resetFields = () => {
    setCurrentPassword("");
    setNewPassword("");
    setConfirmNewPassword("");
  };

  const handlePasswordUpdate = async (
    event: React.FormEvent<HTMLFormElement>,
  ) => {
    event.preventDefault();

    if (newPassword !== confirmNewPassword) {
      toast.error("New passwords do not match.");
      return;
    }

    try {
      await AuthService.updatePassword(currentPassword, newPassword);
      resetFields();
      toast.success("Password updated successfully.");
    } catch (error) {
      if (error instanceof Error) {
        toast.error(`Failed to update password. ${error.message}`);
      } else {
        toast.error("Failed to update password.");
      }
      console.error("Error:", error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="p-6 bg-white shadow-md rounded w-full max-w-md mx-auto">
        <div className="mb-4">
          <h2 className="text-2xl font-semibold">Account Settings</h2>
          <div className="mt-2">
            <strong>Name:</strong> {user.name}
          </div>
          <div>
            <strong>Email:</strong> {user.email}
          </div>
        </div>
        <form onSubmit={handlePasswordUpdate}>
          <h3 className="text-xl font-semibold mb-2">Update Password</h3>
          <input
            type="password"
            value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)}
            placeholder="Current Password"
            className="w-full p-2 mb-3 border rounded"
            required
          />
          <input
            type="password"
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            placeholder="New Password"
            className="w-full p-2 mb-3 border rounded"
            required
          />
          <input
            type="password"
            value={confirmNewPassword}
            onChange={(e) => setConfirmNewPassword(e.target.value)}
            placeholder="Confirm New Password"
            className="w-full p-2 mb-3 border rounded"
            required
          />
          <button
            type="submit"
            className="w-full p-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            disabled={!currentPassword || !newPassword || !confirmNewPassword}
          >
            Update Password
          </button>
        </form>
      </div>
    </div>
  );
};

export default AccountSettings;
