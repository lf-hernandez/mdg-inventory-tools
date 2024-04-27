import { ChangeEvent } from "react";

type Props = {
  label: string;
  name: string;
  value: string | number;
  onChange: (e: ChangeEvent<HTMLInputElement & HTMLTextAreaElement>) => void;
  type?: string;
  required?: boolean;
  multiline?: boolean;
  disabled?: boolean;
};

export const InputField = ({
  label,
  name,
  value,
  onChange,
  type = "text",
  required = false,
  multiline = false,
  disabled = false,
}: Props) => {
  const inputClasses = `
  focus:ring-blue-500
  focus:border-blue-500
  w-full
  shadow-sm
  sm:text-sm
  border
  border-gray-300
  rounded-md
  py-2
  px-3
  disabled:bg-gray-100
  disabled:text-gray-500
  disabled:border-none
`;
  return (
    <div>
      <label
        htmlFor={name}
        className="block text-sm font-medium text-gray-700 mb-1"
      >
        {label}
      </label>
      {multiline ? (
        <textarea
          disabled={disabled}
          rows={5}
          name={name}
          id={name}
          value={value}
          onChange={onChange}
          required={required}
          className={inputClasses}
        />
      ) : (
        <input
          disabled={disabled}
          type={type}
          name={name}
          id={name}
          value={value}
          onChange={onChange}
          required={required}
          className={inputClasses}
        />
      )}
    </div>
  );
};
