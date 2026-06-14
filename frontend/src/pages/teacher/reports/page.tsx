import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
} from '@mui/material';
import {
  Download as DownloadIcon,
  PictureAsPdf as PdfIcon,
  Assessment as AssessmentIcon,
  People as PeopleIcon,
  TrendingUp as TrendingUpIcon,
} from '@mui/icons-material';

const TeacherReports = () => {
  const reportTypes = [
    {
      title: 'Class Performance Report',
      description: 'Overall class performance metrics and statistics',
      icon: <AssessmentIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Student Progress Report',
      description: 'Individual student progress and achievement tracking',
      icon: <TrendingUpIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Attendance Summary',
      description: 'Attendance records and patterns for all classes',
      icon: <PeopleIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
    {
      title: 'Assessment Results',
      description: 'Detailed analysis of quiz and exam results',
      icon: <PdfIcon sx={{ fontSize: 48 }} />,
      action: 'Generate Report',
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Reports
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Generate and view reports for your classes and students
        </Typography>
      </Box>

      <Grid container spacing={3}>
        {reportTypes.map((report, index) => (
          <Grid size={{ xs: 12, md: 6 }} key={index}>
            <Card variant="outlined">
              <CardContent>
                <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                  <Box sx={{ mr: 2, color: 'primary.main' }}>
                    {report.icon}
                  </Box>
                  <Box>
                    <Typography variant="h6">{report.title}</Typography>
                    <Typography variant="body2" color="text.secondary">
                      {report.description}
                    </Typography>
                  </Box>
                </Box>
                <Button
                  variant="outlined"
                  startIcon={<DownloadIcon />}
                  fullWidth
                >
                  {report.action}
                </Button>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Card sx={{ mt: 3 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Recent Reports
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Your recently generated reports will appear here.
          </Typography>
        </CardContent>
      </Card>
    </Container>
  );
};

export default TeacherReports;
