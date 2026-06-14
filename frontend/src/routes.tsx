import React from "react";
import { Navigate, Route, Routes } from "react-router-dom";

import { leftMenuBottomItems, leftMenuItems } from "@/menu-items";
import AppLayout from "@/pages/app/layout";
import AuthLayout from "@/pages/auth/layout";
import Loading from "@/pages/loading.tsx";
import NotFound from "@/pages/not-found";
import { MenuItem } from "@/types/types";
import { ProtectedRoute } from "@/components/auth/ProtectedRoute";

// Statically import all possible pages for build
const modules = import.meta.glob("./pages/**/page.tsx");
const newModules = import.meta.glob("./pages/**/new.tsx");
const idModules = import.meta.glob("./pages/**/[id].tsx");
const editModules = import.meta.glob("./pages/**/[id]/edit/page.tsx");

// Merge all modules
const allModules = { ...modules, ...newModules, ...idModules, ...editModules };

// Lazy load page components
const lazyLoad = (path: string) => {
  // Handle different paths based on the route
  let key: string;
  
  // Remove /dashboard prefix for file path mapping since files don't have it in their directory structure
  const filePath = path.replace(/^\/dashboard/, "");
  
  if (path === "/") {
    key = "./pages/auth/sign-in/page.tsx"; // Root is now sign-in page
  } else if (path === "/landing") {
    key = "./pages/page.tsx"; // Landing page moved to /landing
  } else if (path === "/dashboard") {
    key = "./pages/app/dashboard/page.tsx"; // Dashboard page
  } else if (path === "/sign-in" || path === "/sign-up" || path === "/password-reset" ||
    path === "/password-sent" || path === "/password-new" ||
    path === "/get-verification" || path === "/set-verification" ||
    path === "/terms-and-conditions" || path === "/privacy-policy") {
    key = `./pages/auth${path}/page.tsx`; // Auth pages at root level
  } else if (filePath.includes("/new")) {
    // Handle new/create pages - they are named new.tsx not page.tsx
    key = `./pages/app${filePath}.tsx`;
  } else if (filePath.includes("/:id")) {
    // Handle dynamic routes - check both patterns
    const dynamicPath = filePath.replace("/:id", "/[id]");
    const nestedDynamicPath = filePath.replace("/:id", "/[id]/edit");
    
    // Try nested edit pattern first
    if (allModules[`./pages/app${nestedDynamicPath}/page.tsx`]) {
      key = `./pages/app${nestedDynamicPath}/page.tsx`;
    } else if (allModules[`./pages/app${dynamicPath}/page.tsx`]) {
      key = `./pages/app${dynamicPath}/page.tsx`;
    } else if (allModules[`./pages/app${dynamicPath}.tsx`]) {
      key = `./pages/app${dynamicPath}.tsx`;
    } else {
      key = `./pages/app${filePath}/page.tsx`; // Default fallback
    }
  } else {
    key = `./pages/app${filePath}/page.tsx`;
  }

  const importer = allModules[key];

  // If file not found fallback to 404
  if (!importer) {
    return <Navigate to="/404" replace />;
  }

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

    // Add route for current item with role protection
    const routeElement = lazyLoad(item.href);
    const protectedElement = item.roles ? (
      <ProtectedRoute allowedRoles={item.roles}>{routeElement}</ProtectedRoute>
    ) : (
      routeElement
    );
    routes.push(<Route key={item.id} path={item.href} element={protectedElement} />);

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

// Manual routes for dynamic forms (CRUD operations)
const formRoutes = [
  <Route key="subject-categories-new" path="/dashboard/academic-foundation/subject-categories/new" element={lazyLoad("/dashboard/academic-foundation/subject-categories/new")} />,
  <Route key="subject-categories-edit" path="/dashboard/academic-foundation/subject-categories/:id" element={lazyLoad("/dashboard/academic-foundation/subject-categories/:id")} />,
  <Route key="academic-years-new" path="/dashboard/academic-foundation/academic-years/new" element={lazyLoad("/dashboard/academic-foundation/academic-years/new")} />,
  <Route key="academic-years-edit" path="/dashboard/academic-foundation/academic-years/:id" element={lazyLoad("/dashboard/academic-foundation/academic-years/:id")} />,
  <Route key="semesters-new" path="/dashboard/academic-foundation/semesters/new" element={lazyLoad("/dashboard/academic-foundation/semesters/new")} />,
  <Route key="semesters-edit" path="/dashboard/academic-foundation/semesters/:id" element={lazyLoad("/dashboard/academic-foundation/semesters/:id")} />,
  <Route key="subjects-new" path="/dashboard/curriculum/subjects/new" element={lazyLoad("/dashboard/curriculum/subjects/new")} />,
  <Route key="subjects-edit" path="/dashboard/curriculum/subjects/:id" element={lazyLoad("/dashboard/curriculum/subjects/:id")} />,
  <Route key="phases-new" path="/dashboard/curriculum/phases/new" element={lazyLoad("/dashboard/curriculum/phases/new")} />,
  <Route key="phases-edit" path="/dashboard/curriculum/phases/:id" element={lazyLoad("/dashboard/curriculum/phases/:id")} />,
  <Route key="elements-new" path="/dashboard/curriculum/elements/new" element={lazyLoad("/dashboard/curriculum/elements/new")} />,
  <Route key="elements-edit" path="/dashboard/curriculum/elements/:id" element={lazyLoad("/dashboard/curriculum/elements/:id")} />,
  <Route key="subelements-new" path="/dashboard/curriculum/subelements/new" element={lazyLoad("/dashboard/curriculum/subelements/new")} />,
  <Route key="subelements-edit" path="/dashboard/curriculum/subelements/:id" element={lazyLoad("/dashboard/curriculum/subelements/:id")} />,
  <Route key="reports-publish" path="/dashboard/reports/:id/publish" element={lazyLoad("/dashboard/reports/:id/publish")} />,
];

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
        {/* Manual form routes for CRUD operations */}
        {formRoutes}
      </Route>

      {/* 404 route */}
      <Route path="/404" element={<NotFound />} />
      <Route path="*" element={<Navigate to="/404" replace />} />
    </Routes>
  );
};

export default AppRoutes;
