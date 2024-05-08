import { FormEvent, useState } from "react";

type Props = {
  onSubmit: ({
    newPassword,
    confirmedNewPassword,
    currentPassword,
  }: {
    newPassword: string;
    confirmedNewPassword: string;
    currentPassword: string;
  }) => void;
  onBack: () => void;
};

export const UpdatePasswordForm = ({ onSubmit, onBack }: Props) => {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmedNewPassword, setConfirmedNewPassword] = useState("");

  const isSubmitDisabled =
    !currentPassword ||
    !newPassword ||
    !confirmedNewPassword ||
    newPassword !== confirmedNewPassword;

  const resetFields = () => {
    setCurrentPassword("");
    setNewPassword("");
    setConfirmedNewPassword("");
  };

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    onSubmit({ newPassword, confirmedNewPassword, currentPassword });
    resetFields();
  }
  return (
    <div>
      <button className="flex mb-4 w-full" onClick={onBack}>
        <div className="mr-2">&larr;</div>
        <p>Back</p>
      </button>
      <form onSubmit={handleSubmit}>
        <h3 className="text-xl font-semibold mb-2">Update Password</h3>
        <input
          type="password"
          value={currentPassword}
          onChange={(e) => setCurrentPassword(e.target.value)}
          placeholder="Current Password"
          className="w-full p-2 mb-3 border rounded"
          required
          autoComplete="current-password"
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
          value={confirmedNewPassword}
          onChange={(e) => setConfirmedNewPassword(e.target.value)}
          placeholder="Confirm New Password"
          className="w-full p-2 mb-3 border rounded"
          required
          autoComplete="new-password"
        />
        <button
          type="submit"
          className={`w-full p-2 rounded ${
            isSubmitDisabled
              ? "bg-gray-300 text-gray-600 cursor-not-allowed"
              : "bg-blue-500 text-white hover:bg-blue-600 "
          }`}
          disabled={isSubmitDisabled}
        >
          Update Password
        </button>
      </form>
    </div>
  );
};
