/**
 * NUSA Education Dashboard
 * Main dashboard for teachers showing curriculum, assessment, and student achievement metrics
 */

import DashboardNusaStats from "./sections/dashboard-nusa-stats";
import DashboardNusaCurriculumProgress from "./sections/dashboard-nusa-curriculum-progress";
import DashboardNusaAssessmentAnalytics from "./sections/dashboard-nusa-assessment-analytics";
import DashboardNusaStudentAchievement from "./sections/dashboard-nusa-student-achievement";
import DashboardNusaAIGeneration from "./sections/dashboard-nusa-ai-generation";
import DashboardNusaTaskOverview from "./sections/dashboard-nusa-task-overview";
import dayjs, { Dayjs } from "dayjs";
import weekday from "dayjs/plugin/weekday";
import { useState } from "react";
import { Link as RouterLink } from "react-router-dom";

import { Breadcrumbs, Button, FormControl, Tooltip, Typography, Link } from "@mui/material";
import { Grid } from "@mui/material";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateRange, DateRangePicker } from "@mui/x-date-pickers-pro";

dayjs.extend(weekday);

export default function NusaDashboardPage() {
  const [dateRange, setDateRange] = useState<DateRange<Dayjs>>([dayjs().weekday(-7), dayjs().weekday(-1)]);
  
  return (
    <Grid container spacing={5}>
      <Grid container spacing={2.5} className="w-full" size={12}>
        <Grid size={{ xs: 12, md: "grow" }}>
          <Typography variant="h1" component="h1" className="mb-0">
            Dashboard Guru NUSA
          </Typography>
          <Breadcrumbs>
            <Link component={RouterLink} to="/dashboard" color="inherit">
              Dashboard
            </Link>
            <Typography variant="body2">Ringkasan Pendidikan</Typography>
          </Breadcrumbs>
        </Grid>

        <Grid size={{ xs: 12, md: "auto" }} className="flex flex-row items-start gap-2">
          <FormControl variant="standard" className="surface-standard mb-0 w-full md:w-auto">
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DateRangePicker
                slotProps={{
                  textField: { size: "small", variant: "standard" },
                  desktopPaper: { className: "outlined" },
                }}
                value={dateRange}
                onChange={(newValue) => setDateRange(newValue)}
              />
            </LocalizationProvider>
          </FormControl>

          <Tooltip title="Lihat Detail Tugas">
            <Button
              className="surface-standard flex-none"
              size="medium"
              color="primary"
              variant="surface"
              component={RouterLink}
              to="/tasks"
            >
              Tugas Saya
            </Button>
          </Tooltip>
        </Grid>
      </Grid>

      <Grid container size={12} className="items-start">
        <Grid container size={{ lg: 6, xs: 12 }} className="items-start">
          <Grid size={12}>
            <DashboardNusaStats />
          </Grid>
          <Grid size={12}>
            <DashboardNusaCurriculumProgress />
          </Grid>
        </Grid>
        <Grid size={{ lg: 6, xs: 12 }}>
          <DashboardNusaAssessmentAnalytics />
        </Grid>
      </Grid>

      <Grid container size={12}>
        <DashboardNusaStudentAchievement />
        <DashboardNusaAIGeneration />
        <DashboardNusaTaskOverview />
      </Grid>
    </Grid>
  );
}
