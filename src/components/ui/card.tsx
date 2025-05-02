import React from "react";

export const Card = ({ children, className = "" }: any) => (
  <div className={`rounded-lg border shadow-sm ${className}`}>{children}</div>
);

export const CardContent = ({ children }: any) => (
  <div className="p-4">{children}</div>
);
