import { toast } from "react-hot-toast";
import { useCurrentUser } from "../hooks/useCurrentUser";
import { AuthService } from "../services/AuthService";
import { InputField } from "../components/InputField";
import { useState } from "react";
import { UpdatePasswordForm } from "../components/UpdatePasswordForm";

const AccountSettings = () => {
  const { user } = useCurrentUser();
  const [showUpdatePasswordForm, setShowUpdatePasswordForm] = useState(false);

  const handlePasswordUpdate = async ({
    newPassword,
    confirmedNewPassword,
    currentPassword,
  }: {
    newPassword: string;
    confirmedNewPassword: string;
    currentPassword: string;
  }) => {
    if (newPassword !== confirmedNewPassword) {
      toast.error("New passwords do not match.");
      return;
    }

    try {
      await AuthService.updatePassword(currentPassword, newPassword);
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
        {showUpdatePasswordForm ? (
          <UpdatePasswordForm
            onSubmit={({
              newPassword,
              confirmedNewPassword,
              currentPassword,
            }) => {
              handlePasswordUpdate({
                newPassword,
                confirmedNewPassword,
                currentPassword,
              });
            }}
            onBack={() => {
              setShowUpdatePasswordForm(false);
            }}
          />
        ) : (
          <div className="mb-4">
            <h2 className="text-2xl font-semibold">Account Settings</h2>
            <div className="my-4">
              <InputField
                disabled
                label="Name"
                value={user.name ?? ""}
                name="name"
                onChange={() => {}}
              />
            </div>
            <div className="mb-4">
              <InputField
                disabled
                label="Email"
                value={user.email ?? ""}
                name="email"
                onChange={() => {}}
              />
            </div>
            <div className="mb-4">
              <InputField
                disabled
                label="Role"
                value={user.role ?? ""}
                name="role"
                onChange={() => {}}
              />
            </div>
            <hr />
            <button
              className="flex justify-between my-4 w-full"
              onClick={() => {
                setShowUpdatePasswordForm(true);
              }}
            >
              <p>Update password</p>
              <div>&rarr;</div>
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default AccountSettings;
