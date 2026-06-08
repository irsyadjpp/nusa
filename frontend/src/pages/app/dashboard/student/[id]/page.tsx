/**
 * Student Progress Page
 * Individual student progress view
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  Card,
  CardContent,
  CircularProgress,
  Alert,
  TextField,
} from '@mui/material';
import {
  ArrowBack,
  TrendingUp,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { getStudentAchievement, getStudentProgress } from '@/api/achievement';
import AchievementCard from '@/components/achievement/AchievementCard';
import CompetencyProgress from '@/components/achievement/CompetencyProgress';
import StudentTrajectory from '@/components/achievement/StudentTrajectory';

const StudentProgressPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [studentData, setStudentData] = useState<any>(null);
  const [progressData, setProgressData] = useState<any[]>([]);
  const [trajectoryData, setTrajectoryData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedPeriod, setSelectedPeriod] = useState<string>('current');

  useEffect(() => {
    if (id) {
      loadStudentData(id);
    }
  }, [id, selectedPeriod]);

  const loadStudentData = async (studentId: string) => {
    setLoading(true);
    setError(null);
    try {
      const [achievementData, progressResult, trajectoryResult] = await Promise.all([
        getStudentAchievement(studentId, selectedPeriod),
        getStudentProgress(studentId),
        fetch(`/api/achievements/${studentId}/trajectory`).then(res => res.json()),
      ]);
      setStudentData(achievementData);
      setProgressData(progressResult);
      setTrajectoryData(trajectoryResult);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data siswa');
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Alert severity="error">
        {error}
      </Alert>
    );
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/dashboard')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Progress Siswa</Typography>
        </Box>
        <TextField
          select
          label="Periode"
          value={selectedPeriod}
          onChange={(e) => setSelectedPeriod(e.target.value)}
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="current">Semester Ini</MenuItem>
          <MenuItem value="last">Semester Lalu</MenuItem>
          <MenuItem value="all">Semua</MenuItem>
        </TextField>
      </Box>

      <Grid container spacing={3}>
        <Grid item xs={12} md={4}>
          {studentData && (
            <AchievementCard
              achievement={studentData}
              onClick={() => {}}
            />
          )}
        </Grid>

        <Grid item xs={12} md={8}>
          {trajectoryData && (
            <StudentTrajectory
              student_id={id || ''}
              student_name={studentData?.student_name || 'Siswa'}
              trajectory_points={trajectoryData.trajectory_points || []}
            />
          )}
        </Grid>

        <Grid item xs={12}>
          <Typography variant="h6" gutterBottom>
            Progress Kompetensi
          </Typography>
          <Grid container spacing={3}>
            {progressData.map((progress, index) => (
              <Grid item xs={12} sm={6} md={4} key={index}>
                <CompetencyProgress
                  progress={progress}
                  showDetails
                />
              </Grid>
            ))}
          </Grid>
        </Grid>
      </Grid>
    </Box>
  );
};

export default StudentProgressPage;
