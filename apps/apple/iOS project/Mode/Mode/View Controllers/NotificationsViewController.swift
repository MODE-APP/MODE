//
//  NotificationsViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/17/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class NotificationsViewController: UIViewController {

    // MARK: - Outlets
    
    @IBOutlet weak var notificationsCollectionView: UITableView!
    
    // MARK: - Properties
    
    var loadingNotifications: Bool = false
    var loadedAllNotifications: Bool = false
    var dataSource: Int = 0
    var numberOfNotifications: Int = 60
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        dataSource = 5
        
        
        notificationsCollectionView.delegate = self
        notificationsCollectionView.dataSource = self
    }
    
    
    // MARK: - Custom Functions
    
    func getImages (completion: @escaping () -> Void) {
        let timer = Timer(timeInterval: 0.7, repeats: false) { (_) in
            completion()
        }
        RunLoop.current.add(timer, forMode: .common)
    }
    
    func loadMoreNotifications() {
        getImages() {
            DispatchQueue.main.async {
                if self.loadedAllNotifications == false {
                    self.dataSource += 15
                }
                if self.dataSource >= self.numberOfNotifications {
                    self.loadedAllNotifications = true
                }
                self.notificationsCollectionView.reloadData()
                self.loadingNotifications = false
            }
        }
    }

    func scrollViewDidScroll(_ scrollView: UIScrollView) {
        let heightFromBottom = scrollView.contentSize.height - scrollView.contentOffset.y
        if heightFromBottom < 1300 && loadingNotifications == false && loadedAllNotifications == false{
            loadMoreNotifications()
            loadingNotifications = true
        }
    }
    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */

} // End of class

extension NotificationsViewController: UITableViewDelegate, UITableViewDataSource {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        if dataSource == numberOfNotifications {
            return dataSource
        } else {
            return dataSource + 1
        }
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        if indexPath.row == dataSource && !loadedAllNotifications {
            if let cell = notificationsCollectionView.dequeueReusableCell(withIdentifier: "loadingCell") as? activityIndicatorCell {
                cell.activityIndicatorView.startAnimating()
                return cell
            } else {
                return UITableViewCell()
            }
        } else if let cell = notificationsCollectionView.dequeueReusableCell(withIdentifier: "notificationCell", for: indexPath) as? notificationTableViewCell {
            return cell
        } else {
            return UITableViewCell()
        }
        
    }
    
//    func tableView(_ tableView: UITableView, viewForFooterInSection section: Int) -> UIView? {
//        let view = LoadingAdditionContentIndicatiorUIView()
//        view.configureViews()
//        return view
//    }
//
//    func tableView(_ tableView: UITableView, heightForFooterInSection section: Int) -> CGFloat {
//        return 55
//    }
}
