import { NavLink } from 'react-router-dom'
import styles from './Navbar.module.css'
import ThemeControls from '../ThemeControls/ThemeControls'

const Navbar = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.navLinks}>
        <NavLink 
          to="/" 
          className={({ isActive }) => 
            isActive ? styles.activeLink : styles.link
          }
          end
        >
          Home
        </NavLink>
        <NavLink 
          to="/experience" 
          className={({ isActive }) => 
            isActive ? styles.activeLink : styles.link
          }
        >
          Credentials
        </NavLink>
      </div>
      <div className={styles.themeToggle}>
        <ThemeControls />
      </div>
    </nav>
  )
}

export default Navbar 