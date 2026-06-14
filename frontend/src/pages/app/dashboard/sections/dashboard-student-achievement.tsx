/**
 * Dashboard Student Achievement Component
 * Line chart showing student competency progress
 */

import { Card, CardContent, Typography, Box, Skeleton } from "@mui/material";

export default function DashboardStudentAchievement() {
  return (
    <>
      <Typography variant="h6" component="h6" className="mt-2 mb-3">
        Progres Kompetensi Siswa
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
