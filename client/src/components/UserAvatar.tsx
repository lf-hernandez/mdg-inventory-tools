type Props = {
  fullName: string;
};

export const UserAvatar = ({ fullName }: Props) => {
  const getInitials = (name: string) => {
    const nameArray = name.split(" ");
    const initials = nameArray
      .map((word) => word.charAt(0))
      .slice(0, 2)
      .join("");
    return initials.toUpperCase();
  };

  return (
    <div className="w-10 h-10 flex items-center justify-center rounded-full bg-blue-500 text-white">
      {getInitials(fullName)}
    </div>
  );
};
