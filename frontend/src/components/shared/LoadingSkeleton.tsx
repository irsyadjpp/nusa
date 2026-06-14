/**
 * Loading Skeleton Component
 * Reusable loading skeleton for better UX during data fetching
 */

import { Box, Skeleton, Card, CardContent } from '@mui/material';

interface LoadingSkeletonProps {
  variant?: 'list' | 'detail' | 'card' | 'text';
  count?: number;
}

export const LoadingSkeleton: React.FC<LoadingSkeletonProps> = ({
  variant = 'list',
  count = 3
}) => {
  const renderListSkeleton = () => (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
      {Array.from({ length: count }).map((_, index) => (
        <Box key={index} sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Skeleton variant="circular" width={40} height={40} />
          <Box sx={{ flex: 1 }}>
            <Skeleton variant="text" width="60%" height={24} />
            <Skeleton variant="text" width="40%" height={20} />
          </Box>
          <Skeleton variant="rectangular" width={100} height={32} />
        </Box>
      ))}
    </Box>
  );

  const renderDetailSkeleton = () => (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
      <Box>
        <Skeleton variant="text" width="40%" height={32} />
        <Skeleton variant="text" width="80%" height={24} sx={{ mt: 1 }} />
      </Box>
      <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(2, 1fr)', gap: 2 }}>
        {Array.from({ length: 4 }).map((_, index) => (
          <Box key={index}>
            <Skeleton variant="text" width="30%" height={20} />
            <Skeleton variant="text" width="60%" height={24} sx={{ mt: 0.5 }} />
          </Box>
        ))}
      </Box>
      <Box>
        <Skeleton variant="text" width="20%" height={20} />
        <Skeleton variant="text" width="100%" height={80} sx={{ mt: 0.5 }} />
      </Box>
    </Box>
  );

  const renderCardSkeleton = () => (
    <Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))', gap: 2 }}>
      {Array.from({ length: count }).map((_, index) => (
        <Card key={index}>
          <CardContent>
            <Skeleton variant="text" width="60%" height={24} />
            <Skeleton variant="text" width="100%" height={16} sx={{ mt: 2 }} />
            <Skeleton variant="text" width="80%" height={16} sx={{ mt: 1 }} />
            <Skeleton variant="text" width="80%" height={16} sx={{ mt: 1 }} />
            <Box sx={{ mt: 2, display: 'flex', justifyContent: 'space-between' }}>
              <Skeleton variant="rectangular" width={60} height={24} />
              <Skeleton variant="rectangular" width={80} height={24} />
            </Box>
          </CardContent>
        </Card>
      ))}
    </Box>
  );

  const renderTextSkeleton = () => (
    <Box>
      {Array.from({ length: count }).map((_, index) => (
        <Skeleton key={index} variant="text" width={index % 3 === 0 ? '100%' : index % 3 === 1 ? '80%' : '60%'} height={20} sx={{ mb: 1 }} />
      ))}
    </Box>
  );

  switch (variant) {
    case 'list':
      return renderListSkeleton();
    case 'detail':
      return renderDetailSkeleton();
    case 'card':
      return renderCardSkeleton();
    case 'text':
      return renderTextSkeleton();
    default:
      return renderListSkeleton();
  }
};