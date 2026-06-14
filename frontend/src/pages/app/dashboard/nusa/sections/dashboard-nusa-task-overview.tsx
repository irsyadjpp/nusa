/**
 * Dashboard NUSA Task Overview
 * Shows pending tasks and recent activities for teachers
 */

import { Box, Card, CardContent, Typography, List, ListItem, ListItemText, ListItemIcon, Chip, Button } from "@mui/material";
import AssignmentIcon from "@mui/icons-material/Assignment";
import RateReviewIcon from "@mui/icons-material/RateReview";
import DescriptionIcon from "@mui/icons-material/Description";
import EventIcon from "@mui/icons-material/Event";
import ArrowForwardIcon from "@mui/icons-material/ArrowForward";

export default function DashboardNusaTaskOverview() {
  const pendingTasks = [
    { id: 1, title: "Review TP Matematika Kelas 7", type: "review", priority: "high", due: "Hari ini" },
    { id: 2, title: "Selesaikan Evaluasi Asesmen IPA", type: "evaluation", priority: "medium", due: "Besok" },
    { id: 3, title: "Generate Modul Ajar Bahasa Indonesia", type: "generation", priority: "low", due: "3 hari" },
    { id: 4, title: "Buat Rapor Naratif Siswa", type: "report", priority: "high", due: "5 hari" },
  ];

  const getTaskIcon = (type: string) => {
    switch (type) {
      case "review":
        return <RateReviewIcon color="primary" />;
      case "evaluation":
        return <AssignmentIcon color="secondary" />;
      case "generation":
        return <DescriptionIcon color="primary" />;
      case "report":
        return <DescriptionIcon color="success" />;
      default:
        return <AssignmentIcon />;
    }
  };

  const getPriorityColor = (priority: string) => {
    switch (priority) {
      case "high":
        return "error";
      case "medium":
        return "warning";
      case "low":
        return "info";
      default:
        return "default";
    }
  };

  const getPriorityLabel = (priority: string) => {
    switch (priority) {
      case "high":
        return "Tinggi";
      case "medium":
        return "Sedang";
      case "low":
        return "Rendah";
      default:
        return priority;
    }
  };

  return (
    <>
      <Box sx={{ width: { lg: '66.67%', xs: '100%' } }}>
        <Typography variant="h6" component="h6" className="mb-3">
          Tugas Pending
        </Typography>

        <Card className="h-64">
          <CardContent>
            <List>
              {pendingTasks.map((task) => (
                <ListItem
                  key={task.id}
                  sx={{
                    borderBottom: 1,
                    borderColor: "divider",
                    py: 2,
                  }}
                >
                  <ListItemIcon>
                    {getTaskIcon(task.type)}
                  </ListItemIcon>
                  <ListItemText
                    primary={task.title}
                    secondary={
                      <Box sx={{ display: "flex", alignItems: "center", gap: 1, mt: 1 }}>
                        <Chip
                          label={getPriorityLabel(task.priority)}
                          size="small"
                          color={getPriorityColor(task.priority) as any}
                        />
                        <Typography variant="caption" color="text.secondary">
                          <EventIcon sx={{ fontSize: 14, mr: 0.5 }} />
                          {task.due}
                        </Typography>
                      </Box>
                    }
                  />
                </ListItem>
              ))}
            </List>
          </CardContent>
        </Card>
      </Box>

      <Box sx={{ width: { lg: '33.33%', xs: '100%' } }}>
        <Typography variant="h6" component="h6" className="mb-3">
          Aktivitas Cepat
        </Typography>

        <Card className="h-64">
          <CardContent>
            <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
              <Button
                variant="outlined"
                fullWidth
                startIcon={<DescriptionIcon />}
                component="a"
                href="/tp/create"
              >
                Buat TP Baru
              </Button>
              <Button
                variant="outlined"
                fullWidth
                startIcon={<AssignmentIcon />}
                component="a"
                href="/assessment/create"
              >
                Buat Asesmen
              </Button>
              <Button
                variant="outlined"
                fullWidth
                startIcon={<RateReviewIcon />}
                component="a"
                href="/reports"
              >
                Lihat Rapor
              </Button>
              <Button
                variant="contained"
                fullWidth
                startIcon={<ArrowForwardIcon />}
                component="a"
                href="/tasks"
              >
                Lihat Semua Tugas
              </Button>
            </Box>
          </CardContent>
        </Card>
      </Box>
    </>
  );
}
