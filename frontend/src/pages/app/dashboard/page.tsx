import { Grid, Typography } from "@mui/material";

import DashboardStats from "./sections/dashboard-stats";
import DashboardCurriculumProgress from "./sections/dashboard-curriculum-progress";
import DashboardStudentAchievement from "./sections/dashboard-student-achievement";
import DashboardQuickActions from "./sections/dashboard-quick-actions";
import DashboardRecentActivity from "./sections/dashboard-recent-activity";

export default function DashboardPage() {
  return (
    <Grid container spacing={5}>
      <Grid container spacing={2.5} className="w-full" size={12}>
        <Grid size={{ xs: 12, md: "grow" }}>
          <Typography variant="h1" component="h1" className="mb-0">
            Dashboard Guru NUSA
          </Typography>
        </Grid>
      </Grid>

      <Grid container size={12} className="items-start">
        <Grid container size={{ lg: 6, xs: 12 }} className="items-start">
          <Grid size={12}>
            <DashboardStats />
          </Grid>
          <Grid size={12}>
            <DashboardCurriculumProgress />
          </Grid>
        </Grid>
        <Grid size={{ lg: 6, xs: 12 }}>
          <DashboardStudentAchievement />
        </Grid>
      </Grid>

      <Grid container size={12}>
        <Grid size={{ lg: 6, xs: 12 }}>
          <DashboardQuickActions />
        </Grid>
        <Grid size={{ lg: 6, xs: 12 }}>
          <DashboardRecentActivity />
        </Grid>
      </Grid>
    </Grid>
  );
}
