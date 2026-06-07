import React from "react";
import { Navigate, Route, Routes } from "react-router-dom";

import { leftMenuBottomItems, leftMenuItems } from "@/menu-items";
import AppLayout from "@/pages/app/layout";
import AuthLayout from "@/pages/auth/layout";
import Loading from "@/pages/loading.tsx";
import NotFound from "@/pages/not-found";
import { MenuItem } from "@/types/types";

// Statically import all possible pages for build
const modules = import.meta.glob("./pages/**/page.tsx");

// Lazy load page components
const lazyLoad = (path: string) => {
  // Handle different paths based on the route
  let key: string;
  if (path === "/") {
    key = "./pages/auth/sign-in/page.tsx"; // Root is now sign-in page
  } else if (path === "/landing") {
    key = "./pages/page.tsx"; // Landing page moved to /landing
  } else if (path === "/sign-in" || path === "/sign-up" || path === "/password-reset" ||
    path === "/password-sent" || path === "/password-new" ||
    path === "/get-verification" || path === "/set-verification" ||
    path === "/terms-and-conditions" || path === "/privacy-policy") {
    key = `./pages/auth${path}/page.tsx`; // Auth pages at root level
  } else {
    key = `./pages/app${path}/page.tsx`;
  }

  const importer = modules[key];

  // If file not found fallback to 404
  if (!importer) return <Navigate to="/404" replace />;

  const Component = React.lazy(importer as () => Promise<{ default: React.ComponentType<any> }>);

  return (
    <React.Suspense fallback={<Loading />}>
      <Component />
    </React.Suspense>
  );
};

// Recursively generate routes from menu items
const generateRoutesFromMenuItems = (menuItems: MenuItem[]): React.ReactElement[] => {
  return menuItems.flatMap((item: MenuItem) => {
    const routes: React.ReactElement[] = [];

    // Skip external links
    if (item.isExternalLink || !item.href) {
      return [];
    }

    // Add route for current item
    routes.push(<Route key={item.id} path={item.href} element={lazyLoad(item.href)} />);

    // Add routes for children
    if (item.children && item.children.length > 0) {
      routes.push(...generateRoutesFromMenuItems(item.children));
    }

    return routes;
  });
};

// Generate auth routes (moved to root level)
const generateAuthRoutes = (): React.ReactElement[] => {
  return [
    <Route key="sign-in" path="/sign-in" element={lazyLoad("/sign-in")} />,
    <Route key="sign-up" path="/sign-up" element={lazyLoad("/sign-up")} />,
    <Route key="password-reset" path="/password-reset" element={lazyLoad("/password-reset")} />,
    <Route key="password-sent" path="/password-sent" element={lazyLoad("/password-sent")} />,
    <Route key="password-new" path="/password-new" element={lazyLoad("/password-new")} />,
    <Route key="get-verification" path="/get-verification" element={lazyLoad("/get-verification")} />,
    <Route key="set-verification" path="/set-verification" element={lazyLoad("/set-verification")} />,
    <Route key="terms-and-conditions" path="/terms-and-conditions" element={lazyLoad("/terms-and-conditions")} />,
    <Route key="privacy-policy" path="/privacy-policy" element={lazyLoad("/privacy-policy")} />,
  ];
};

// Generate routes from both menu arrays
const mainRoutes = generateRoutesFromMenuItems(leftMenuItems);
const bottomRoutes = generateRoutesFromMenuItems(leftMenuBottomItems);
const authRoutes = generateAuthRoutes();

// Main Routes component
const AppRoutes = () => {
  return (
    <Routes>
      {/* Login page at root */}
      <Route path="/" element={lazyLoad("/")} />
      {/* Landing page moved to /landing */}
      <Route path="/landing" element={lazyLoad("/landing")} />
      {/* Auth routes with AuthLayout at root level */}
      <Route element={<AuthLayout />}>
        {authRoutes}
      </Route>
      {/* App routes with AppLayout */}
      <Route element={<AppLayout />}>
        {/* Routes generated from menu items */}
        {mainRoutes}
        {bottomRoutes}
      </Route>

      {/* 404 route */}
      <Route path="/404" element={<NotFound />} />
      <Route path="*" element={<Navigate to="/404" replace />} />
    </Routes>
  );
};

export default AppRoutes;
