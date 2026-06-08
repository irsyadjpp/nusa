// Shared utilities
// This directory will contain utility functions used across the application

import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function isPathMatch(path: string, pattern: string): boolean {
  // Remove trailing slashes
  const cleanPath = path.replace(/\/$/, "");
  const cleanPattern = pattern.replace(/\/$/, "");

  // Exact match
  if (cleanPath === cleanPattern) {
    return true;
  }

  // Pattern match (e.g., /dashboard matches /dashboard/settings)
  if (cleanPath.startsWith(cleanPattern + "/")) {
    return true;
  }

  return false;
}
