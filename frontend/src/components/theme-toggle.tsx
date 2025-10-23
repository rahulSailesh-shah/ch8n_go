import * as React from "react";
import { MoonIcon, SunIcon } from "lucide-react";

import { useTheme } from "@/components/theme-provider";
import { Switch } from "@/components/ui/switch";

export function ThemeToggle() {
  const { theme, setTheme } = useTheme();
  const isDarkMode = theme === "dark";

  const toggleTheme = () => {
    setTheme(isDarkMode ? "light" : "dark");
  };

  return (
    <div className="flex items-center gap-2">
      <SunIcon className="h-[1.2rem] w-[1.2rem]" />
      <Switch 
        checked={isDarkMode}
        onCheckedChange={toggleTheme}
        aria-label="Toggle theme"
      />
      <MoonIcon className="h-[1.2rem] w-[1.2rem]" />
    </div>
  );
}
