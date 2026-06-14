/**
 * Error Boundary Component
 * Catches JavaScript errors in child components and displays a fallback UI
 */

import { Component, ErrorInfo, ReactNode } from 'react';
import { Box, Typography, Button, Card, CardContent, Alert } from '@mui/material';
import { Refresh as RefreshIcon } from '@mui/icons-material';

interface ErrorBoundaryProps {
  children: ReactNode;
  fallback?: ReactNode;
}

interface ErrorBoundaryState {
  hasError: boolean;
  error: Error | null;
  errorInfo: ErrorInfo | null;
}

export class ErrorBoundary extends Component<ErrorBoundaryProps, ErrorBoundaryState> {
  constructor(props: ErrorBoundaryProps) {
    super(props);
    this.state = {
      hasError: false,
      error: null,
      errorInfo: null,
    };
  }

  static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    return {
      hasError: true,
      error,
      errorInfo: null,
    };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    this.setState({
      error,
      errorInfo,
    });

    // Log error to console for debugging
    console.error('ErrorBoundary caught an error:', error, errorInfo);

    // In production, you might want to send this to an error tracking service
    // logErrorToService(error, errorInfo);
  }

  handleReset = () => {
    this.setState({
      hasError: false,
      error: null,
      errorInfo: null,
    });
  };

  render() {
    if (this.state.hasError) {
      // Custom fallback UI
      if (this.props.fallback) {
        return this.props.fallback;
      }

      // Default fallback UI
      return (
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '400px', p: 3 }}>
          <Card sx={{ maxWidth: 600, width: '100%' }}>
            <CardContent>
              <Alert severity="error" sx={{ mb: 2 }}>
                Terjadi kesalahan tidak terduga
              </Alert>
              
              <Typography variant="h6" gutterBottom>
                Ups! Ada yang salah
              </Typography>
              
              <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                Terjadi kesalahan saat memuat halaman ini. Silakan coba lagi atau hubungi administrator jika masalah berlanjut.
              </Typography>

              {import.meta.env.DEV && this.state.error && (
                <Box sx={{ mt: 2, p: 2, bgcolor: 'grey.100', borderRadius: 1, overflow: 'auto' }}>
                  <Typography variant="caption" component="pre" sx={{ color: 'error.main' }}>
                    {this.state.error.toString()}
                  </Typography>
                  {this.state.errorInfo && (
                    <Typography variant="caption" component="pre" sx={{ mt: 1, color: 'text.secondary' }}>
                      {this.state.errorInfo.componentStack}
                    </Typography>
                  )}
                </Box>
              )}

              <Button
                variant="contained"
                startIcon={<RefreshIcon />}
                onClick={this.handleReset}
                sx={{ mt: 2 }}
              >
                Coba Lagi
              </Button>
            </CardContent>
          </Card>
        </Box>
      );
    }

    return this.props.children;
  }
}