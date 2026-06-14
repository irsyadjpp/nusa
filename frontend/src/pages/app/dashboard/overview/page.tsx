/**
 * Dashboard Overview Page
 * Main dashboard with progress overview
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  TrendingUp,
  People,
  Assignment,
  Star,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { getClassAchievement } from '@/api/achievement';

const DashboardOverviewPage: React.FC = () => {
  const navigate = useNavigate();
  const [classData, setClassData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    loadDashboardData();
  }, []);

  const loadDashboardData = async () => {
    setLoading(true);
    setError(null);
    try {
      // Load class achievement data
      const data = await getClassAchievement('class-id-placeholder');
      setClassData(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data dashboard');
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

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Dashboard Progress
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3, mb: 3 }}>
        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card
            sx={{ cursor: 'pointer' }}
            onClick={() => navigate('/tp')}
          >
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <Assignment color="primary" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="body2" color="text.secondary">
                    Total TP
                  </Typography>
                  <Typography variant="h4">12</Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card
            sx={{ cursor: 'pointer' }}
            onClick={() => navigate('/assessment')}
          >
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <TrendingUp color="success" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="body2" color="text.secondary">
                    Asesmen
                  </Typography>
                  <Typography variant="h4">8</Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card
            sx={{ cursor: 'pointer' }}
            onClick={() => navigate('/evidence')}
          >
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <People color="info" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="body2" color="text.secondary">
                    Bukti
                  </Typography>
                  <Typography variant="h4">45</Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>

        <Box sx={{ width: { xs: '100%', sm: '50%', md: '25%' } }}>
          <Card
            sx={{ cursor: 'pointer' }}
            onClick={() => navigate('/reports')}
          >
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <Star color="warning" sx={{ fontSize: 40 }} />
                <Box>
                  <Typography variant="body2" color="text.secondary">
                    Rapor
                  </Typography>
                  <Typography variant="h4">3</Typography>
                </Box>
              </Box>
            </CardContent>
          </Card>
        </Box>
      </Box>

      {classData && (
        <Card sx={{ mt: 3 }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Ringkasan Pencapaian Kelas
            </Typography>
            <Typography variant="body2">
              Data pencapaian kelas dimuat. Komponen ClassAchievementSummary dinonaktifkan sementara.
            </Typography>
            <pre>{JSON.stringify(classData, null, 2)}</pre>
          </CardContent>
        </Card>
      )}
    </Box>
  );
};

export default DashboardOverviewPage;
