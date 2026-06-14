/**
 * Dashboard NUSA Curriculum Progress
 * Shows curriculum implementation progress and TP generation metrics
 */

import { Card, CardContent, Typography, Box, Skeleton } from "@mui/material";

export default function DashboardNusaCurriculumProgress() {
  return (
    <>
      <Typography variant="h6" component="h6" className="mb-3">
        Progres Kurikulum
      </Typography>

      <Card className="h-80">
        <CardContent>
          <Box sx={{ height: 265, display: 'flex', alignItems: 'center', justifyContent: 'center', flexDirection: 'column', gap: 2 }}>
            <Typography variant="body1" color="text.secondary">
              Chart loading...
            </Typography>
            <Skeleton variant="rectangular" width="100%" height={200} />
          </Box>
        </CardContent>
      </Card>
    </>
  );
}