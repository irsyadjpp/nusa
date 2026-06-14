/**
 * Dashboard NUSA Assessment Analytics
 * Shows assessment creation, completion, and evaluation metrics
 */

import { Box, Card, CardContent, List, ListItem, Typography, Chip } from "@mui/material";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import PendingIcon from "@mui/icons-material/Pending";
import ScheduleIcon from "@mui/icons-material/Schedule";

export default function DashboardNusaAssessmentAnalytics() {
  const assessmentData = [
    { id: 1, name: "Asesmen Formatif Matematika", subject: "Matematika", status: "completed", students: 32, evaluated: 30 },
    { id: 2, name: "Asesmen Sumatif Bahasa Indonesia", subject: "Bahasa Indonesia", status: "pending", students: 32, evaluated: 28 },
    { id: 3, name: "Asesmen Diagnostik IPA", subject: "IPA", status: "scheduled", students: 32, evaluated: 0 },
    { id: 4, name: "Asesmen Portofolio Seni", subject: "Seni Budaya", status: "completed", students: 32, evaluated: 32 },
    { id: 5, name: "Asesmen Proyek PKN", subject: "PKN", status: "pending", students: 32, evaluated: 25 },
  ];

  const getStatusIcon = (status: string) => {
    switch (status) {
      case "completed":
        return <CheckCircleIcon color="success" />;
      case "pending":
        return <PendingIcon color="warning" />;
      case "scheduled":
        return <ScheduleIcon color="info" />;
      default:
        return <ScheduleIcon color="info" />;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "completed":
        return "success";
      case "pending":
        return "warning";
      case "scheduled":
        return "info";
      default:
        return "default";
    }
  };

  const getStatusLabel = (status: string) => {
    switch (status) {
      case "completed":
        return "Selesai";
      case "pending":
        return "Dalam Proses";
      case "scheduled":
        return "Terjadwal";
      default:
        return status;
    }
  };

  return (
    <>
      <Typography variant="h6" component="h6" className="mb-3">
        Analitik Asesmen
      </Typography>

      <Card className="h-80">
        <CardContent>
          <Box sx={{ maxHeight: 265, overflowY: "auto" }}>
            <List>
              {assessmentData.map((assessment) => (
                <ListItem
                  key={assessment.id}
                  sx={{
                    borderBottom: 1,
                    borderColor: "divider",
                    py: 2,
                  }}
                >
                  <Box sx={{ flexGrow: 1 }}>
                    <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 1 }}>
                      {getStatusIcon(assessment.status)}
                      <Typography variant="subtitle2" fontWeight="medium">
                        {assessment.name}
                      </Typography>
                    </Box>
                    <Box sx={{ display: "flex", alignItems: "center", gap: 2, mt: 1 }}>
                      <Typography variant="caption" color="text.secondary">
                        {assessment.subject}
                      </Typography>
                      <Chip
                        label={getStatusLabel(assessment.status)}
                        size="small"
                        color={getStatusColor(assessment.status) as any}
                      />
                      <Typography variant="caption" color="text.secondary">
                        {assessment.evaluated}/{assessment.students} dievaluasi
                      </Typography>
                    </Box>
                  </Box>
                </ListItem>
              ))}
            </List>
          </Box>
        </CardContent>
      </Card>
    </>
  );
}
