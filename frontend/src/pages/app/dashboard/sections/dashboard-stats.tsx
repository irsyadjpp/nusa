/**
 * Dashboard Stats Component
 * Key metrics cards for NUSA education platform
 * Following analytics example pattern
 */

import { Link } from "react-router-dom";

import { Box, Card, CardContent, Grid, Typography } from "@mui/material";

import SchoolIcon from "@mui/icons-material/School";
import AssignmentIcon from "@mui/icons-material/Assignment";
import PeopleIcon from "@mui/icons-material/People";
import TrendingUpIcon from "@mui/icons-material/TrendingUp";

export default function DashboardStats() {
  return (
    <>
      <Grid size={{ xs: 12 }}>
        <Typography variant="h6" component="h6" className="mb-3">
          Statistik Utama
        </Typography>

        <Grid container size={12} spacing={2.5} className="flex-none">
          <Grid size={{ xs: 6, sm: 3 }}>
            <Card component={Link} to="/tp" className="flex flex-col p-1 transition-transform hover:scale-[1.02]">
              <Box className="bg-primary-light/10 flex h-28 w-full flex-none items-center justify-center rounded-2xl">
                <AssignmentIcon className="text-primary" sx={{ fontSize: 40 }} />
              </Box>
              <CardContent className="text-center">
                <Typography variant="body1" className="text-text-secondary leading-5 transition-colors">
                  Rencana Pelaksanaan
                </Typography>
                <Typography variant="h5" className="text-leading-5">
                  24
                </Typography>
              </CardContent>
            </Card>
          </Grid>
          <Grid size={{ xs: 6, sm: 3 }}>
            <Card component={Link} to="/assessment" className="flex flex-col p-1 transition-transform hover:scale-[1.02]">
              <Box className="bg-secondary-light/10 flex h-28 w-full flex-none items-center justify-center rounded-2xl">
                <SchoolIcon className="text-secondary" sx={{ fontSize: 40 }} />
              </Box>
              <CardContent className="text-center">
                <Typography variant="body1" className="text-text-secondary leading-5 transition-colors">
                  Asesmen
                </Typography>
                <Typography variant="h5" className="text-leading-5">
                  18
                </Typography>
              </CardContent>
            </Card>
          </Grid>
          <Grid size={{ xs: 6, sm: 3 }}>
            <Card component={Link} to="/students" className="flex flex-col p-1 transition-transform hover:scale-[1.02]">
              <Box className="bg-primary-light/10 flex h-28 w-full flex-none items-center justify-center rounded-2xl">
                <PeopleIcon className="text-primary" sx={{ fontSize: 40 }} />
              </Box>
              <CardContent className="text-center">
                <Typography variant="body1" className="text-text-secondary leading-5 transition-colors">
                  Siswa
                </Typography>
                <Typography variant="h5" className="text-leading-5">
                  156
                </Typography>
              </CardContent>
            </Card>
          </Grid>
          <Grid size={{ xs: 6, sm: 3 }}>
            <Card component={Link} to="/reports" className="flex flex-col p-1 transition-transform hover:scale-[1.02]">
              <Box className="bg-secondary-light/10 flex h-28 w-full flex-none items-center justify-center rounded-2xl">
                <TrendingUpIcon className="text-secondary" sx={{ fontSize: 40 }} />
              </Box>
              <CardContent className="text-center">
                <Typography variant="body1" className="text-text-secondary leading-5 transition-colors">
                  Tingkat Penyelesaian
                </Typography>
                <Typography variant="h5" className="text-leading-5">
                  78%
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
