import React from 'react';
import { useParams } from 'react-router-dom';
import {
  Box,
  Typography,
  Button,
  Grid,
  CircularProgress,
  Alert,
  Card,
  CardContent,
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  LinearProgress,
  Container,
  Breadcrumbs,
  Link,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useNavigate } from 'react-router-dom';
import { useStudentAchievement, useStudentProgress } from '@/services/queries/AchievementQueryService';

const StudentDashboard: React.FC = () => {
  const navigate = useNavigate();
  const { id: studentId } = useParams<{ id: string }>();

  const {
    data: studentAchievements,
    isLoading: isLoadingAchievement,
    error: achievementError,
  } = useStudentAchievement(studentId || '');

  const {
    data: studentProgress,
    isLoading: isLoadingProgress,
    error: progressError,
  } = useStudentProgress(studentId || '');

  if (isLoadingAchievement || isLoadingProgress) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (achievementError || progressError) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Alert severity="error">{achievementError?.message || progressError?.message || 'Failed to load student data'}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
        <Link underline="hover" onClick={() => navigate('/achievement')}>
          Achievement
        </Link>
        <Typography color="text.primary">Student Dashboard</Typography>
      </Breadcrumbs>

      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" component="h1">
          Student Dashboard - {studentId}
        </Typography>
        <Button
          variant="outlined"
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/achievement')}
        >
          Back to Achievement
        </Button>
      </Box>

      <Grid container spacing={3}>
        <Grid size={{ xs: 12 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Student Achievement
              </Typography>
              {studentAchievements && studentAchievements.length > 0 ? (
                <TableContainer>
                  <Table>
                    <TableHead>
                      <TableRow>
                        <TableCell>Competency</TableCell>
                        <TableCell>Mastery Level</TableCell>
                        <TableCell>Score</TableCell>
                        <TableCell>Percentage</TableCell>
                        <TableCell>Achieved Criteria</TableCell>
                        <TableCell>Pending Criteria</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {studentAchievements.map((achievement: any) => (
                        <TableRow key={achievement.competency_id}>
                          <TableCell>{achievement.competency_name}</TableCell>
                          <TableCell>
                            <Chip
                              label={achievement.mastery_level}
                              color={
                                achievement.mastery_level === 'EXEMPLARY'
                                  ? 'success'
                                  : achievement.mastery_level === 'PROFICIENT'
                                  ? 'primary'
                                  : achievement.mastery_level === 'DEVELOPING'
                                  ? 'warning'
                                  : 'error'
                              }
                              size="small"
                            />
                          </TableCell>
                          <TableCell>
                            {achievement.score} / {achievement.max_score}
                          </TableCell>
                          <TableCell>
                            <LinearProgress
                              variant="determinate"
                              value={achievement.percentage}
                              sx={{ minWidth: 100 }}
                            />
                            <Typography variant="caption">{achievement.percentage}%</Typography>
                          </TableCell>
                          <TableCell>
                            {achievement.achieved_criteria.length > 0 ? (
                              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {achievement.achieved_criteria.map((criteria: any) => (
                                  <Chip key={criteria} label={criteria} size="small" color="success" />
                                ))}
                              </Box>
                            ) : (
                              <Typography color="text.secondary">None</Typography>
                            )}
                          </TableCell>
                          <TableCell>
                            {achievement.pending_criteria.length > 0 ? (
                              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {achievement.pending_criteria.map((criteria: any) => (
                                  <Chip key={criteria} label={criteria} size="small" color="warning" />
                                ))}
                              </Box>
                            ) : (
                              <Typography color="text.secondary">None</Typography>
                            )}
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Typography color="text.secondary">No achievement data available</Typography>
              )}
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Competency Progress
              </Typography>
              {studentProgress && studentProgress.length > 0 ? (
                <TableContainer>
                  <Table>
                    <TableHead>
                      <TableRow>
                        <TableCell>Competency</TableCell>
                        <TableCell>Progress</TableCell>
                        <TableCell>Average Score</TableCell>
                        <TableCell>Completed Assessments</TableCell>
                        <TableCell>Total Assessments</TableCell>
                        <TableCell>Mastery Level</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {studentProgress.map((progress: any) => (
                        <TableRow key={progress.competency_id}>
                          <TableCell>{progress.competency_name}</TableCell>
                          <TableCell>
                            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                              <LinearProgress
                                variant="determinate"
                                value={progress.progress_percentage}
                                sx={{ flex: 1, minWidth: 100 }}
                              />
                              <Typography variant="body2">{progress.progress_percentage}%</Typography>
                            </Box>
                          </TableCell>
                          <TableCell>{progress.average_score}</TableCell>
                          <TableCell>
                            {progress.completed_assessments} / {progress.total_assessments}
                          </TableCell>
                          <TableCell>{progress.total_assessments}</TableCell>
                          <TableCell>
                            <Chip
                              label={progress.mastery_level}
                              color={
                                progress.mastery_level === 'EXEMPLARY'
                                  ? 'success'
                                  : progress.mastery_level === 'PROFICIENT'
                                  ? 'primary'
                                  : progress.mastery_level === 'DEVELOPING'
                                  ? 'warning'
                                  : 'error'
                              }
                              size="small"
                            />
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              ) : (
                <Typography color="text.secondary">No progress data available</Typography>
              )}
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Actions
              </Typography>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <Button
                  variant="contained"
                  onClick={() => navigate('/reports/generate', { state: { studentId, classId: null } })}
                >
                  Generate Narrative Report
                </Button>
                <Button variant="outlined" onClick={() => navigate('/achievement')}>
                  Back to Achievement Dashboard
                </Button>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default StudentDashboard;
