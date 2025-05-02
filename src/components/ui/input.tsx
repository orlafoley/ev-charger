import React from "react";

export const Input = ({ className = "", ...props }: any) => (
  <input className={`border rounded p-2 ${className}`} {...props} />
);
