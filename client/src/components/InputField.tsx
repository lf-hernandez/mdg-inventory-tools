import { ChangeEvent } from "react";

type Props = {
  label: string;
  name: string;
  type?: string;
  value: string | number;
  onChange: (e: ChangeEvent<HTMLInputElement & HTMLTextAreaElement>) => void;
  required?: boolean;
  multiline?: boolean;
};

export const InputField = ({
  label,
  name,
  type = "text",
  value,
  onChange,
  required = false,
  multiline = false,
}: Props) => (
  <div>
    <label
      htmlFor={name}
      className="block text-sm font-medium text-gray-700 mb-1"
    >
      {label}
    </label>
    {multiline ? (
      <textarea
        rows={5}
        name={name}
        id={name}
        value={value}
        onChange={onChange}
        required={required}
        className="focus:ring-blue-500 focus:border-blue-500 block w-full shadow-sm sm:text-sm border border-gray-300 rounded-md py-2 px-3 text-base leading-tight"
      />
    ) : (
      <input
        type={type}
        name={name}
        id={name}
        value={value}
        onChange={onChange}
        required={required}
        className="focus:ring-blue-500 focus:border-blue-500 block w-full shadow-sm sm:text-sm border border-gray-300 rounded-md py-2 px-3 text-base leading-tight"
      />
    )}
  </div>
);
