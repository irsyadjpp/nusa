import {
  Box,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
  LinearProgress,
} from '@mui/material';
import {
  Assignment as AssignmentIcon,
  Quiz as QuizIcon,
  School as SchoolIcon,
  TrendingUp as TrendingUpIcon,
} from '@mui/icons-material';

const StudentResults = () => {
  const results = [
    {
      id: '1',
      title: 'Mathematics Chapter 4 Quiz',
      subject: 'Mathematics',
      class: '10-A',
      type: 'Quiz',
      date: '2024-01-10',
      score: 85,
      maxScore: 100,
      grade: 'A',
    },
    {
      id: '2',
      title: 'Science Lab Report',
      subject: 'Science',
      class: '10-A',
      type: 'Assignment',
      date: '2024-01-08',
      score: 78,
      maxScore: 100,
      grade: 'B+',
    },
    {
      id: '3',
      title: 'Physics Chapter 3 Test',
      subject: 'Physics',
      class: '10-A',
      type: 'Exam',
      date: '2024-01-05',
      score: 92,
      maxScore: 100,
      grade: 'A',
    },
    {
      id: '4',
      title: 'English Essay',
      subject: 'English',
      class: '10-A',
      type: 'Assignment',
      date: '2024-01-03',
      score: 88,
      maxScore: 100,
      grade: 'A-',
    },
  ];

  const getTypeIcon = (type: string) => {
    switch (type) {
      case 'Quiz':
        return <QuizIcon />;
      case 'Exam':
        return <SchoolIcon />;
      default:
        return <AssignmentIcon />;
    }
  };

  const getGradeColor = (score: number) => {
    if (score >= 90) return 'success';
    if (score >= 80) return 'primary';
    if (score >= 70) return 'info';
    if (score >= 60) return 'warning';
    return 'error';
  };

  const averageScore = Math.round(results.reduce((acc, r) => acc + r.score, 0) / results.length);

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Results
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View your assessment results and grades
        </Typography>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Average Score
              </Typography>
              <Typography variant="h4" color="primary.main">
                {averageScore}%
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Total Assessments
              </Typography>
              <Typography variant="h4">
                {results.length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Highest Score
              </Typography>
              <Typography variant="h4" color="success.main">
                {Math.max(...results.map(r => r.score))}%
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Recent Trend
              </Typography>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
                <TrendingUpIcon color="success" />
                <Typography variant="h4" color="success.main">
                  +5%
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Grid container spacing={3}>
        {results.map((result) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={result.id}>
            <Card variant="outlined">
              <CardContent>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                    <Box sx={{ color: 'primary.main' }}>
                      {getTypeIcon(result.type)}
                    </Box>
                    <Typography variant="subtitle1" fontWeight="medium">
                      {result.title}
                    </Typography>
                  </Box>
                  <Chip
                    label={result.grade}
                    size="small"
                    color={getGradeColor(result.score) as any}
                  />
                </Box>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  {result.subject} • Class {result.class}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  {result.date}
                </Typography>
                <Box sx={{ mt: 2 }}>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 0.5 }}>
                    <Typography variant="body2" color="text.secondary">
                      Score
                    </Typography>
                    <Typography variant="body2" fontWeight="medium">
                      {result.score}/{result.maxScore}
                    </Typography>
                  </Box>
                  <LinearProgress
                    variant="determinate"
                    value={result.score}
                    sx={{
                      height: 8,
                      borderRadius: 4,
                      backgroundColor: 'grey.200',
                      '& .MuiLinearProgress-bar': {
                        borderRadius: 4,
                        backgroundColor: getGradeColor(result.score) === 'success' ? '#4caf50' :
                                       getGradeColor(result.score) === 'primary' ? '#2196f3' :
                                       getGradeColor(result.score) === 'info' ? '#00bcd4' :
                                       getGradeColor(result.score) === 'warning' ? '#ff9800' : '#f44336',
                      },
                    }}
                  />
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default StudentResults;
