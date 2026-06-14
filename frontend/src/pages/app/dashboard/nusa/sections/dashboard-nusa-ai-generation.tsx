/**
 * Dashboard NUSA AI Generation
 * Shows AI assistant generation statistics and approval rates
 */

import { Box, Card, CardContent, Typography, Chip, CircularProgress } from "@mui/material";
import SmartToyIcon from "@mui/icons-material/SmartToy";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import AutoAwesomeIcon from "@mui/icons-material/AutoAwesome";

export default function DashboardNusaAIGeneration() {
  const aiStats = {
    totalGenerated: 45,
    approved: 38,
    pending: 5,
    rejected: 2,
    approvalRate: 84.4,
  };

  const recentGenerations = [
    { type: "TP", subject: "Matematika", status: "approved", time: "2 jam yang lalu" },
    { type: "Modul Ajar", subject: "Bahasa Indonesia", status: "pending", time: "4 jam yang lalu" },
    { type: "Rubrik", subject: "IPA", status: "approved", time: "6 jam yang lalu" },
    { type: "TP", subject: "Seni Budaya", status: "approved", time: "1 hari yang lalu" },
  ];

  const getStatusColor = (status: string) => {
    switch (status) {
      case "approved":
        return "success";
      case "pending":
        return "warning";
      case "rejected":
        return "error";
      default:
        return "default";
    }
  };

  const getStatusLabel = (status: string) => {
    switch (status) {
      case "approved":
        return "Disetujui";
      case "pending":
        return "Menunggu";
      case "rejected":
        return "Ditolak";
      default:
        return status;
    }
  };

  return (
    <>
      <Box sx={{ width: { lg: '50%', xs: '100%' } }}>
        <Typography variant="h6" component="h6" className="mb-3">
          Statistik Generasi AI
        </Typography>

        <Card>
          <CardContent>
            <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between", mb: 3 }}>
              <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
                <SmartToyIcon color="primary" sx={{ fontSize: 32 }} />
                <Box>
                  <Typography variant="h5" fontWeight="medium">
                    {aiStats.totalGenerated}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Total Generasi
                  </Typography>
                </Box>
              </Box>
              <Box sx={{ position: "relative", display: "inline-flex" }}>
                <CircularProgress
                  variant="determinate"
                  value={aiStats.approvalRate}
                  size={80}
                  color="success"
                />
                <Box
                  sx={{
                    top: 0,
                    left: 0,
                    bottom: 0,
                    right: 0,
                    position: "absolute",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <Typography variant="body2" fontWeight="medium">
                    {aiStats.approvalRate}%
                  </Typography>
                </Box>
              </Box>
            </Box>

            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2 }}>
              <Box sx={{ width: { xs: '33.33%' }, textAlign: "center" }}>
                <Box sx={{ textAlign: "center" }}>
                  <CheckCircleIcon color="success" sx={{ fontSize: 24, mb: 1 }} />
                  <Typography variant="h6" fontWeight="medium">
                    {aiStats.approved}
                  </Typography>
                  <Typography variant="caption" color="text.secondary">
                    Disetujui
                  </Typography>
                </Box>
              </Box>
              <Box sx={{ width: { xs: '33.33%' }, textAlign: "center" }}>
                <Box sx={{ textAlign: "center" }}>
                  <AutoAwesomeIcon color="warning" sx={{ fontSize: 24, mb: 1 }} />
                  <Typography variant="h6" fontWeight="medium">
                    {aiStats.pending}
                  </Typography>
                  <Typography variant="caption" color="text.secondary">
                    Menunggu
                  </Typography>
                </Box>
              </Box>
              <Box sx={{ width: { xs: '33.33%' }, textAlign: "center" }}>
                <Box sx={{ textAlign: "center" }}>
                  <Typography variant="h6" fontWeight="medium" color="error">
                    {aiStats.rejected}
                  </Typography>
                  <Typography variant="caption" color="text.secondary">
                    Ditolak
                  </Typography>
                </Box>
              </Box>
            </Box>
          </CardContent>
        </Card>
      </Box>

      <Box sx={{ width: { lg: '50%', xs: '100%' } }}>
        <Typography variant="h6" component="h6" className="mb-3">
          Generasi Terbaru
        </Typography>

        <Card className="h-48">
          <CardContent>
            <Box sx={{ maxHeight: 150, overflowY: "auto" }}>
              {recentGenerations.map((gen, index) => (
                <Box
                  key={index}
                  sx={{
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "space-between",
                    py: 1.5,
                    borderBottom: index < recentGenerations.length - 1 ? 1 : 0,
                    borderColor: "divider",
                  }}
                >
                  <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
                    <SmartToyIcon color="primary" sx={{ fontSize: 20 }} />
                    <Box>
                      <Typography variant="body2" fontWeight="medium">
                        {gen.type} - {gen.subject}
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        {gen.time}
                      </Typography>
                    </Box>
                  </Box>
                  <Chip
                    label={getStatusLabel(gen.status)}
                    size="small"
                    color={getStatusColor(gen.status) as any}
                  />
                </Box>
              ))}
            </Box>
          </CardContent>
        </Card>
      </Box>
    </>
  );
}
