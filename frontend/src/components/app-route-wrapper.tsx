/**
 * App Route Wrapper
 * Handles protected route logic with React Router hooks
 */

import { Suspense } from "react";
import { useLocation } from "react-router-dom";

import Loading from "@/pages/loading";
import { useAuth } from "@/features/auth";
import { ProtectedRoute } from "@/features/auth";

const AppRouteWrapper = ({ children }: { children: React.ReactNode }) => {
  const location = useLocation();
  const { loading } = useAuth();

  // Public routes (no authentication required)
  const publicRoutes = ['/', '/sign-up', '/password-reset', '/password-sent', '/password-new', '/get-verification', '/set-verification'];
  
  const isPublicRoute = publicRoutes.includes(location.pathname);

  return (
    <Suspense fallback={<Loading />}>
      {loading && !isPublicRoute ? (
        <Loading />
      ) : isPublicRoute ? (
        <>{children}</>
      ) : (
        <ProtectedRoute>
          {children}
        </ProtectedRoute>
      )}
    </Suspense>
  );
};

export default AppRouteWrapper;