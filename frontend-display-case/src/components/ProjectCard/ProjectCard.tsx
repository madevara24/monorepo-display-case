import { FC, useEffect, useState } from 'react'
import styles from './ProjectCard.module.css'

interface GitHubStats {
  stars: number;
  forks: number;
  lastUpdated: string;
  isLoading: boolean;
  error: boolean;
}

interface ProjectCardProps {
  name: string;
  repo: string;
  description: string;
  technologies: string[];
}

const ProjectCard: FC<ProjectCardProps> = ({ 
  name, 
  repo, 
  description, 
  technologies
}) => {
  const [stats, setStats] = useState<GitHubStats>({
    stars: 0,
    forks: 0,
    lastUpdated: '',
    isLoading: true,
    error: false
  });

  // Extract GitHub username and repository name from URL
  const getRepoInfo = (url: string) => {
    try {
      const urlParts = new URL(url).pathname.split('/').filter(Boolean);
      return {
        username: urlParts[0],
        repo: urlParts[1]
      };
    } catch {
      return { username: '', repo: '' };
    }
  };

  const repoInfo = getRepoInfo(repo);

  // Fetch GitHub stats with localStorage caching
  useEffect(() => {
    const fetchWithCache = async () => {
      if (!repoInfo.username || !repoInfo.repo) {
        setStats(prev => ({ ...prev, isLoading: false, error: true }));
        return;
      }
      
      const cacheKey = `github-${repoInfo.username}-${repoInfo.repo}`;
      const cachedData = localStorage.getItem(cacheKey);
      
      // Use cache if it exists and is less than 24 hours old
      if (cachedData) {
        try {
          const parsed = JSON.parse(cachedData);
          if (parsed.timestamp > Date.now() - 24 * 60 * 60 * 1000) {
            setStats({
              ...parsed.data,
              isLoading: false,
              error: false
            });
            return;
          }
        } catch (e) {
          // Invalid cache, continue to fetch
          console.error('Error parsing cached GitHub data:', e);
        }
      }
      
      try {
        const response = await fetch(
          `https://api.github.com/repos/${repoInfo.username}/${repoInfo.repo}`
        );
        
        if (!response.ok) {
          throw new Error(`Failed to fetch repository data: ${response.status}`);
        }
        
        const data = await response.json();
        const lastUpdated = formatRelativeTime(new Date(data.updated_at));
        
        const statsData = {
          stars: data.stargazers_count,
          forks: data.forks_count,
          lastUpdated
        };
        
        // Save to cache
        localStorage.setItem(cacheKey, JSON.stringify({
          timestamp: Date.now(),
          data: statsData
        }));
        
        setStats({
          ...statsData,
          isLoading: false,
          error: false
        });
      } catch (error) {
        console.error('Error fetching GitHub stats:', error);
        setStats(prev => ({ ...prev, isLoading: false, error: true }));
      }
    };
    
    fetchWithCache();
  }, [repoInfo.username, repoInfo.repo]);

  // Helper function to format date to relative time
  const formatRelativeTime = (date: Date): string => {
    const now = new Date();
    const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000);
    
    const intervals = {
      year: 31536000,
      month: 2592000,
      week: 604800,
      day: 86400,
      hour: 3600,
      minute: 60
    };

    for (const [unit, seconds] of Object.entries(intervals)) {
      const interval = Math.floor(diffInSeconds / seconds);
      if (interval >= 1) {
        return interval === 1 ? `1 ${unit} ago` : `${interval} ${unit}s ago`;
      }
    }
    
    return 'just now';
  };

  return (
    <a href={repo} target="_blank" rel="noopener noreferrer" className={styles.card}>
      <div className={styles.content}>
        <div className={styles.header}>
          {/* GitHub icon */}
          <div className={styles.githubIcon}>
            <svg viewBox="0 0 24 24" width="24" height="24">
              <path d="M12 2C6.477 2 2 6.477 2 12c0 4.42 2.865 8.163 6.839 9.489.5.092.682-.217.682-.48 0-.237-.008-.866-.013-1.7-2.782.603-3.369-1.34-3.369-1.34-.454-1.156-1.11-1.462-1.11-1.462-.908-.62.069-.608.069-.608 1.003.07 1.531 1.03 1.531 1.03.892 1.529 2.341 1.087 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.11-4.555-4.943 0-1.091.39-1.984 1.029-2.683-.103-.253-.446-1.27.098-2.647 0 0 .84-.268 2.75 1.026A9.578 9.578 0 0112 6.836c.85.004 1.705.114 2.504.336 1.909-1.294 2.747-1.026 2.747-1.026.546 1.377.203 2.394.1 2.647.64.699 1.028 1.592 1.028 2.683 0 3.842-2.339 4.687-4.566 4.934.359.309.678.92.678 1.855 0 1.337-.012 2.415-.012 2.743 0 .267.18.578.688.48C19.138 20.16 22 16.416 22 12c0-5.523-4.477-10-10-10z"/>
            </svg>
          </div>
          
          {/* Name and repo info */}
          <div className={styles.titleContainer}>
            <h3 className={styles.name}>{name}</h3>
            <span className={styles.repoInfo}>
              {repoInfo.username}/{repoInfo.repo}
            </span>
          </div>
        </div>
        
        <p className={styles.description}>{description}</p>
        
        {/* Tech stack section */}
        <div className={styles.technologies}>
          {technologies.map((tech, index) => (
            <span key={index} className={styles.tech}>
              {tech}
            </span>
          ))}
        </div>
        
        {/* GitHub stats with loading state */}
        <div className={styles.stats}>
          {stats.isLoading ? (
            <span className={styles.stat}>Loading stats...</span>
          ) : stats.error ? (
            <span className={styles.stat}>Stats unavailable</span>
          ) : (
            <>
              <span className={styles.stat}>
                <svg viewBox="0 0 16 16" width="16" height="16">
                  <path d="M8 .25a.75.75 0 01.673.418l1.882 3.815 4.21.612a.75.75 0 01.416 1.279l-3.046 2.97.719 4.192a.75.75 0 01-1.088.791L8 12.347l-3.766 1.98a.75.75 0 01-1.088-.79l.72-4.194L.818 6.374a.75.75 0 01.416-1.28l4.21-.611L7.327.668A.75.75 0 018 .25z"/>
                </svg>
                {stats.stars}
              </span>
              <span className={styles.stat}>
                <svg viewBox="0 0 16 16" width="16" height="16">
                  <path d="M5 3.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm0 2.122a2.25 2.25 0 10-1.5 0v.878A2.25 2.25 0 005.75 8.5h1.5v2.128a2.251 2.251 0 101.5 0V8.5h1.5a2.25 2.25 0 002.25-2.25v-.878a2.25 2.25 0 10-1.5 0v.878a.75.75 0 01-.75.75h-4.5A.75.75 0 015 6.25v-.878zm3.75 7.378a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm3-8.75a.75.75 0 100-1.5.75.75 0 000 1.5z"/>
                </svg>
                {stats.forks}
              </span>
              <span className={styles.stat}>
                <svg viewBox="0 0 16 16" width="16" height="16" className={styles.updatedIcon}>
                  <path d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.477l4.817 4.816a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" />
                </svg>
                {stats.lastUpdated}
              </span>
            </>
          )}
        </div>
      </div>
    </a>
  )
}

export default ProjectCard 